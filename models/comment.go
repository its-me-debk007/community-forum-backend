package models

import "time"

type Comment struct {
	CommentID   uint      `json:"comment_id" gorm:"primaryKey;autoIncrement"`
	CommentMsg  string    `json:"comment_msg"`
	Edited      bool      `json:"edited"`
	CommentedAt time.Time `json:"commented_at"`
	CommentorID uint      `json:"commentor_id"`
	Commentor   User      `json:"commentor" gorm:"foreignKey:CommentorID"`
}
