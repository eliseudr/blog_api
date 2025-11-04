package repository

import (
	"github.com/eliseudr/blog_api/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) GetAll() ([]models.BlogPost, error) {
	var posts []models.BlogPost
	err := r.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) GetByID(id uint) (models.BlogPost, error) {
	var post models.BlogPost
	err := r.db.Preload("Comments").First(&post, id).Error

	if err == gorm.ErrRecordNotFound {
		return models.BlogPost{}, nil
	}
	if err != nil {
		return models.BlogPost{}, err
	}

	if post.Comments == nil {
		post.Comments = []models.Comment{}
	}

	return post, nil
}

func (r *PostRepository) Create(post *models.BlogPost) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) UpdateCommentCount(id uint) error {
	var count int64
	err := r.db.Model(&models.Comment{}).Where("post_id = ?", id).Count(&count).Error
	if err != nil {
		return err
	}
	return r.db.Model(&models.BlogPost{}).Where("id = ?", id).Update("comment_count", count).Error
}
