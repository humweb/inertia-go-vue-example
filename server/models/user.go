package models

import (
	"github.com/jaswdr/faker/v2"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (User) TableName() string {
	return "users"
}

func UsersSeed(db *gorm.DB) error {
	fake := faker.New()
	var users []User

	for i := 0; i < 30; i++ {
		users = append(users, User{FirstName: fake.Person().FirstName(), LastName: fake.Person().LastName(), Email: fake.Internet().Email()})
	}
	return db.Debug().Create(&users).Error

}
