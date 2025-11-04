package models

import "gorm.io/gorm"

// Define the Comment model
type Comment struct {
	gorm.Model
	Content string   `gorm:"not null" json:"content"`
	PostID  uint     `gorm:"not null" json:"post_id"`
	Post    BlogPost `gorm:"foreignKey:PostID" json:"-"`
	// "-" excludes it from JSON serialization, removing the nested Post object
}
