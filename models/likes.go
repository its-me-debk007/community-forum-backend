package models

import "github.com/lib/pq"

type Likes struct {
	ID    uint          `json:"id" gorm:"primaryKey"`
	Users pq.Int64Array `json:"users" gorm:"type:integer[]"`
}
