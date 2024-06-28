package models

var TestUser User

type Achievement struct {
	Id    uint
	Img   string
	Title string
	Link  string
}

type User struct {
	ID           uint
	Name         string
	Email        string
	Password     string
	GitURL       string
	Role         string
	Achievements []Achievement
}

func CreateTestUser(username string, password string) *User {
	achievments := []Achievement{
		{
			1,
			"assets/img1.png",
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
