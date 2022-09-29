package models

import (
	"time"
)

type User struct {
	// gorm.Model
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	ProfilePic string    `json:"profile_pic"`
}
