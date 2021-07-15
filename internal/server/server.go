package server

import (
	"net/http"
	"taskozon/internal/store"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  *store.Store
}

func NewServer(maxSize int) *Server {
	server := &Server{
		router: gin.Default(),
		store:  store.New(maxSize),
	}
	server.ConfigureRouter()
	return server
}

func (s *Server) ConfigureRouter() {
	s.router.POST("/set", SetHandler(s.store))
	s.router.POST("/get", GetHandler(s.store))
	s.router.POST("/del", DeleteHandler(s.store))
	s.router.GET("/getallkeys", GetAllKeysHandler(s.store))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Start(port string) {
	s.router.Run(port)
}
