package author

import (
	"context"
	"errors"
	"fmt"

	"github.com/XATAB1CH/achievement-holder/postgresql"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
}

func (r *repository) Create(ctx context.Context, author *Author) error {
	q := `
		INSERT INTO author (name, age) VALUES ($1, $2) RETURNING id
	`

	p := r.client.QueryRow(ctx, q, author.Name, 123)
	if err := (&p).Scan(&author.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			return newErr
		}
		return err
	}

	return nil
}

func (r *repository) FindAll(ctx context.Context) (u []Author, err error) {
	q := `
		SELECT id, name FROM public.author;
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors := make([]Author, 0)

	for rows.Next() {
		var author Author

		err = rows.Scan(&author.ID, &author.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (Author, error) {
	q := `
		SELECT id, name FROM public.author WHERE id = $1
	`

	var author Author
	p := r.client.QueryRow(ctx, q, id)
	err := (&p).Scan(&author.ID, &author.Name)
	if err != nil {
		return Author{}, err
	}

	return author, nil
}

func (r *repository) Update(ctx context.Context, author Author) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client) Repository {
	return &repository{
		client: client,
	}
}
