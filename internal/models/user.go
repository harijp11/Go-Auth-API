package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"column:id;primaryKey"`
	Username     string    `json:"username" gorm:"column:username;size:100;not null"`
	Email        string    `json:"email" gorm:"column:email;size:255;not null;uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"column:password_hash;not null"`
	PhoneNumber  string    `json:"phone_number" gorm:"column:phone_number;size:20"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
