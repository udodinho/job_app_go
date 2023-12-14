package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" validate:"required,uuid"`
	Name      string    `json:"name" validate:"required,lte=255"`
	Email     string    `gorm:"unique" json:"email" validate:"required,lte=255"`
	Password  string    `json:"password,omitempty" validate:"required,lte=255"`
	Jobs      []Job     `]son:"jobs" gorm:"ForeignKey:CreatedBy"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}


func GetUserByID(id uuid.UUID) (*User, error) {
	var u User

	DB.Where("id = ?", id).First(&u)

	return &u, nil
}

func GetUserByEmail(email string) (User, error) {
	 var u User
	 
	 DB.Where("email=?", email).Find(&u)

	return u, nil
}

func (u *User) CreateUser() (*User, error) {
	
	DB.Create(&u)

	return u, nil
}