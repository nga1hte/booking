package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/nga1hte/booking/db/sqlc"
	"github.com/nga1hte/booking/token"
	"github.com/nga1hte/booking/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.getUsers)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.DELETE("/users/:id", server.deleteUser)
	authRoutes.PUT("/users", server.updateUser)

	authRoutes.POST("/bookings", server.createBooking)
	authRoutes.GET("/bookings", server.getUserBookings)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
