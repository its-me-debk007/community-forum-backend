package models

type Post struct {
	PostID uint `json:"post_id" gorm:"primaryKey"`
	PostName string `json:"post_name"`
	PostDescription string `json:"post_description"`
	AuthorID uint `json:"author_id"`
	Author Author `json:"author" gorm:"foreignKey:AuthorID"`
	LikesCount uint `json:"likes_count"`
	IsLiked bool `json:"is_liked"`
}