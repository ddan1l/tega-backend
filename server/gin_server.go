package server

import (
	"fmt"

	"github.com/ddan1l/tega-api/config"
	"github.com/ddan1l/tega-api/database"

	auth_handler "github.com/ddan1l/tega-api/handlers/auth"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	app  *gin.Engine
	conf *config.Config
	db   database.Database
}

func NewGinServer(conf *config.Config, db database.Database) Server {
	app := gin.Default()

	return &ginServer{
		app:  app,
		db:   db,
		conf: conf,
	}
}

func (s *ginServer) Start() {
	authHandler := auth_handler.NewAuthHandler()

	s.app.POST("/login", authHandler.Login)

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)

	s.app.Run(serverUrl)
}
