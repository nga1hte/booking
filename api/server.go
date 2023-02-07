package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/nga1hte/booking/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.getUsers)
	router.DELETE("/users/:id", server.deleteUser)
	router.PUT("/users", server.updateUser)
	router.POST("/bookings", server.createBooking)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
