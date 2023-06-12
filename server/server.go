package server

import (
	"log"

	"github.com/aabdullahgungor/go-restapi-rabbitmq/constants"

	"github.com/aabdullahgungor/go-restapi-rabbitmq/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   constants.SERVERPORT,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := router.ConfigRoutes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run("localhost:" + s.port))
}
