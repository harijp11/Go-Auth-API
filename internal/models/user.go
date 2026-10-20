package models

import "time"

type User struct {
    ID           uint      `json:"id" gorm:"primaryKey"`
    Username     string    `json:"username" gorm:"size:100;not null"`
    Email        string    `json:"email" gorm:"size:255;not null;uniqueIndex"`
    PasswordHash string    `json:"-" gorm:"not null"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}