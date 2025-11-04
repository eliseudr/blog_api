package server

import (
	"log"
	"net/http"

	"github.com/eliseudr/blog_api/middleware"
	"github.com/eliseudr/blog_api/router"
	"gorm.io/gorm"
)

type BlogAPIServer struct {
	addr string
	db   *gorm.DB
}

func NewBlogAPIServer(addr string, db *gorm.DB) *BlogAPIServer {
	return &BlogAPIServer{addr: addr, db: db}
}

func (s *BlogAPIServer) Run() error {
	routes := router.SetupRoutes(s.db)

	server := http.Server{
		Addr:    s.addr,
		Handler: middleware.ErrorHandler(middleware.Logging(routes)),
	}

	log.Printf("Server is running on port %s", s.addr)

	return server.ListenAndServe()
}
