package models

import "github.com/lib/pq"

type Post struct {
	PostID          uint           `json:"post_id" gorm:"primaryKey;autoIncrement"`
	PostTitle       string         `json:"post_title"`
	PostDescription string         `json:"post_description"`
	PostImages      pq.StringArray `json:"post_images" gorm:"type:text[]"`
	AuthorID        uint           `json:"author_id"`
	Author          User           `json:"author" gorm:"foreignKey:AuthorID"`
	LikesCount      uint           `json:"likes_count"`
	IsLiked         bool           `json:"is_liked"`
	Comments        []Comment      `json:"comments" gorm:"foreignKey:CommentID"`
}