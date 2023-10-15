package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rewired-gh/pet_adoption_app/db/sqlc"
)

// Serves HTTP requests for banking service
type Server struct {
	store  sqlc.Store
	router *gin.Engine
}

// Creates a new server and setup routing
func NewServer(store sqlc.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/user", server.createUser)
	router.GET("/user", server.getUser)
	router.GET("/user/auth", server.authUser)
	router.POST("/user/contact", server.addContact)

	server.router = router
	return server
}

// Start an HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
