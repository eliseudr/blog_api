package repository

import (
	"github.com/eliseudr/blog_api/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) GetByIDWithPost(id uint) (models.Comment, error) {
	var comment models.Comment
	err := r.db.Preload("Post.Comments").First(&comment, id).Error
	return comment, err
}
