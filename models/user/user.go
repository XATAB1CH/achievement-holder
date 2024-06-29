package user

import "github.com/XATAB1CH/achievement-holder/models/achievement"

var TestUser User

type User struct {
	ID           uint
	Name         string
	Email        string
	Password     string
	GitURL       string
	Role         string
	Achievements []achievement.Achievement
}

func CreateTestUser(username string, password string) *User {
	testAchievments := []achievement.Achievement{
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

func (u *User) AddAchievements(achievements ...achievement.Achievement) {
	u.Achievements = append(u.Achievements, achievements...)
}
