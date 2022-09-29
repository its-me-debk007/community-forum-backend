package models

type Post struct {
	PostID          uint   `json:"post_id" gorm:"primaryKey;autoIncrement"`
	PostTitle       string `json:"post_title"`
	PostDescription string `json:"post_description"`
	PostImages      string `json:"post_images"`
	AuthorID        uint   `json:"author_id"`
	Author          User   `json:"author" gorm:"foreignKey:AuthorID"`
	LikesCount      uint   `json:"likes_count"`
	IsLiked         bool   `json:"is_liked"`
}
