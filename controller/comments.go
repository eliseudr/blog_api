package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/eliseudr/blog_api/models"
	"github.com/eliseudr/blog_api/repository"
	"github.com/eliseudr/blog_api/response"
)

type CommentController struct {
	commentRepo *repository.CommentRepository
	postRepo    *repository.PostRepository
}

// Validate all fields for the create comment request.
func (c *CommentController) validateCreateComment(w http.ResponseWriter, r *http.Request) (uint, models.Comment, bool) {
	// Strip url into 2 parts, the post ID and the comments. (in the middle should be the id)
	// Example: Part 1 -> /api/posts/1/comments -> ["api/posts/", "1", "/comments"]
	path := strings.TrimPrefix(r.URL.Path, "/api/posts/")
	parts := strings.Split(path, "/comments")
	if len(parts) != 2 || parts[0] == "" {
		response.Error(w, http.StatusBadRequest, "Invalid URL format")
		return 0, models.Comment{}, false
	}

	postIDStr := parts[0]
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid post ID format")
		return 0, models.Comment{}, false
	}

	post, err := c.postRepo.GetByID(uint(postID))
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch post")
		return 0, models.Comment{}, false
	}

	if post.ID == 0 {
		response.Error(w, http.StatusNotFound, "Post not found")
		return 0, models.Comment{}, false
	}

	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return 0, models.Comment{}, false
	}

	if comment.Content == "" {
		response.Error(w, http.StatusBadRequest, "Content is required")
		return 0, models.Comment{}, false
	}

	return uint(postID), comment, true
}

func NewCommentController(commentRepo *repository.CommentRepository, postRepo *repository.PostRepository) *CommentController {
	return &CommentController{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

// CreateComment adds a new comment to a specific blog post
func (c *CommentController) CreateComment(w http.ResponseWriter, r *http.Request) {
	postID, comment, ok := c.validateCreateComment(w, r)
	if !ok {
		return
	}

	// Insert the post comment into the database.
	comment.PostID = uint(postID)
	if err := c.commentRepo.Create(&comment); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create comment")
		return
	}

	// Update the post comment count.
	if err := c.postRepo.UpdateCommentCount(postID); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update post comment count")
		return
	}

	// Get the created comment with the post.
	createdComment, err := c.commentRepo.GetByIDWithPost(comment.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch created comment")
		return
	}

	// Return the created comment with the post.
	response.Success(w, http.StatusCreated, createdComment)
}
