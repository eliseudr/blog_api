package models

import "gorm.io/gorm"

// Define the BlogPost model
type BlogPost struct {
	gorm.Model
	Title    string    `gorm:"not null"`
	Content  string    `gorm:"not null"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}
