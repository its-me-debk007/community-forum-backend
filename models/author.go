package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name        string `json:"name"`
	ProfilePic string `json:"profile_pic"`
}
