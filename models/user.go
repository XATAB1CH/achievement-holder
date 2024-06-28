package models

var TestUser User

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
	GitURL   string
	Role         string
	Achievements []Achievement
}

func CreateTestUser(username string, password string) *User {
	testAchievments := []Achievement{
		{
			1,
			"Тянка",
			"assets/img1.png",
			"https://google.com",
			1,
		},
		{
			2,
			"Тоже тянка",
			"assets/img1.png",
			"https://google.com",
			1,
		},
	}

	return &User{
		Name:         username,
		Password:     password,
		Achievements: testAchievments,
	}
}

func (u *User) AddAchievements(achievements ...Achievement) {
	u.Achievements = append(u.Achievements, achievements...)
}
