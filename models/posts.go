package models

import "gorm.io/gorm"

// Define the BlogPost model
// CommentCount is used to store the number of comments for the post (Avoiding N+1 problem).
type BlogPost struct {
	gorm.Model
	Title        string    `gorm:"not null" json:"title"`
	Content      string    `gorm:"not null" json:"content"`
	CommentCount int       `gorm:"not null;default:0" json:"comment_count"`
	Comments     []Comment `gorm:"foreignKey:PostID" json:"comments"`
}
