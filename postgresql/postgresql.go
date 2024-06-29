package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/XATAB1CH/achievement-holder/config"
	"github.com/XATAB1CH/achievement-holder/utils"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttempts int, config config.Config) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("host=%s port=%s author=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.Author, config.Password, config.DBName, config.SSLMode)

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		fmt.Println("Failed to connect to PostgreSQL")
	}

	return
}
