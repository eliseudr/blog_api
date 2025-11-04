package router

import (
	"net/http"
	"strings"

	"github.com/eliseudr/blog_api/controller"
	"github.com/eliseudr/blog_api/repository"
	"github.com/eliseudr/blog_api/response"
	"gorm.io/gorm"
)

// SetupRoutes configures and returns a router with all API routes
func SetupRoutes(db *gorm.DB) *http.ServeMux {
	router := http.NewServeMux()

	postRepo := repository.NewPostRepository(db)
	postController := controller.NewPostController(postRepo)

	commentRepo := repository.NewCommentRepository(db)
	commentController := controller.NewCommentController(commentRepo, postRepo)

	router.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			postController.GetPosts(w, r)
		case http.MethodPost:
			postController.CreatePost(w, r)
		default:
			response.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	// One handle for both get and post since they share the same preffix (api/posts/)
	router.HandleFunc("/api/posts/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		// If the URL ends with /comments and the method is POST, create a new comment.
		case strings.HasSuffix(r.URL.Path, "/comments") && r.Method == http.MethodPost:
			commentController.CreateComment(w, r)
		// If the method is GET, get a specific post.
		case r.Method == http.MethodGet:
			postController.GetPost(w, r)
		default:
			response.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	return router
}
