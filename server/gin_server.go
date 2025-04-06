package server

import (
	"fmt"
	"net/http"

	"github.com/ddan1l/tega-api/config"
	"github.com/ddan1l/tega-api/database"
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
	s.app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ww",
		})
	})

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)

	s.app.Run(serverUrl)
}
