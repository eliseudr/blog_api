package models

import "gorm.io/gorm"

// Define the Comment model
type Comment struct {
	gorm.Model
	Content string   `gorm:"not null"`
	PostID  uint     `gorm:"not null"`
	Post    BlogPost `gorm:"foreignKey:PostID"`
}
