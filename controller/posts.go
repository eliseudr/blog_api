package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/eliseudr/blog_api/models"
	"github.com/eliseudr/blog_api/repository"
	"github.com/eliseudr/blog_api/response"
	"gorm.io/gorm"
)

type PostController struct {
	repo *repository.PostRepository
}

type PostListItem struct {
	ID           uint           `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	Title        string         `json:"title"`
	Content      string         `json:"content"`
	CommentCount int            `json:"comment_count"`
}

func NewPostController(repo *repository.PostRepository) *PostController {
	return &PostController{repo: repo}
}

func (c *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	// !! Use only for testing
	// panic("Test 500 error handler")

	posts, err := c.repo.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}

	postList := make([]PostListItem, len(posts))
	for i, post := range posts {
		postList[i] = PostListItem{
			ID:           post.ID,
			CreatedAt:    post.CreatedAt,
			UpdatedAt:    post.UpdatedAt,
			DeletedAt:    post.DeletedAt,
			Title:        post.Title,
			Content:      post.Content,
			CommentCount: post.CommentCount,
		}
	}

	response.Success(w, http.StatusOK, postList)
}

// GetPost retrieves a specific post by ID from the path
func (c *PostController) GetPost(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/posts/")
	if idStr == "" || idStr == r.URL.Path {
		response.Error(w, http.StatusBadRequest, "ID parameter is required")
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID format")
		return
	}

	post, err := c.repo.GetByID(uint(id))
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch post")
		return
	}

	// If the post is not found, return an empty array.
	if post.ID == 0 {
		response.Success(w, http.StatusOK, []models.BlogPost{})
		return
	}

	response.Success(w, http.StatusOK, post)
}

// CreatePost add a new post
func (c *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.BlogPost
	// Decode the request body into the post struct
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if post.Title == "" || post.Content == "" {
		response.Error(w, http.StatusBadRequest, "Title and content are required")
		return
	}

	if err := c.repo.Create(&post); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create post")
		return
	}

	response.Success(w, http.StatusCreated, post)
}
