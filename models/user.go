package models

var TestUser User

type Achievement struct {
	Id    uint   `gorm:"primaryKey" json: "id"`
	Img   string // ссылка на изображение
	Title string // название
	Link  string // ссылка на страницу достижения
}

type User struct {
	ID           uint          `gorm:"primaryKey" json: "id"`
	Name         string        `json: "name"`
	Email        string        `gorm: "unique" json: "email"`
	Password     string        `json: "password"`
	GitURL       string        `json: "giturl"`
	Role         string        `json: "role"`
	Achievements []Achievement `gorm:"many2many:user_achievements" json: "achievements"`
}

func CreateTestUser(username string, password string) *User {
	achievments := []Achievement{
		{
			1,
			"assets/img1.png", // ссылка на изображение из интернета
			"Тянка",
			"https://google.com",
		},
		{
			2,
			"assets/img1.png",
			"Тоже тянка",
			"https://google.com",
		},
	}

	return &User{
		Name:         username,
		Password:     password,
		Achievements: achievments,
	}
}

func (u *User) AddAchievements(achievements ...Achievement) {
	u.Achievements = append(u.Achievements, achievements...)
}
