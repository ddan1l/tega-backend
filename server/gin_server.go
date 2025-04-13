package server

import (
	"fmt"
	"log"
	"os"

	"github.com/ddan1l/tega-backend/config"
	"github.com/ddan1l/tega-backend/database"
	"github.com/ddan1l/tega-backend/factory"
	auth_handler "github.com/ddan1l/tega-backend/handlers/auth"
	auth_middleware "github.com/ddan1l/tega-backend/middleware/auth"
	"github.com/gin-gonic/gin"
)

type ginServer struct {
	app     *gin.Engine
	db      database.Database
	conf    *config.Config
	factory *factory.DefaultFactory
}

func NewGinServer(conf *config.Config, db database.Database) Server {
	app := gin.Default()
	factory := factory.NewDefaultFactory(db)

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return &ginServer{
		app:     app,
		db:      db,
		conf:    conf,
		factory: factory,
	}
}

func (s *ginServer) Start() {
	s.initializeAuthHandler()

	authMiddleware := auth_middleware.NewAuthMiddleware(
		s.factory.CreateAuthUseCase(),
	)

	s.app.Use(authMiddleware.Middleware())

	s.app.GET("/protected", func(c *gin.Context) {
		u, _ := c.Get("user")
		c.JSON(200, gin.H{
			"message": u,
		})
	})

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Run(serverUrl)
}

func (s *ginServer) initializeAuthHandler() {
	authHandler := auth_handler.NewAuthHandler(
		s.factory.CreateAuthUseCase(),
	)

	g := s.app.Group("/auth")

	g.POST("/register", authHandler.Register)
	g.POST("/login", authHandler.Login)
	g.POST("/logout", authHandler.Logout)
}
