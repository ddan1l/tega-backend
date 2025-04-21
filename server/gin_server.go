package server

import (
	"fmt"
	"log"
	"os"

	"github.com/ddan1l/tega-backend/config"
	"github.com/ddan1l/tega-backend/database"
	_ "github.com/ddan1l/tega-backend/docs"
	"github.com/ddan1l/tega-backend/factory"
	auth_handler "github.com/ddan1l/tega-backend/handlers/auth"
	"github.com/ddan1l/tega-backend/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	s.app.Use(middleware.CORSMiddleware())

	if os.Getenv("ENV") != "production" {
		s.app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	authMiddleware := middleware.NewAuthMiddleware(
		s.factory.CreateAuthUseCase(),
	)

	s.initializeAuthHandler(authMiddleware)

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Run(serverUrl)
}

func (s *ginServer) initializeAuthHandler(authMiddleware middleware.AuthMiddleware) {
	authHandler := auth_handler.NewAuthHandler(
		s.factory.CreateAuthUseCase(),
	)

	g := s.app.Group("/auth")

	g.POST("/register", authHandler.Register)
	g.POST("/login", authHandler.Login)
	g.POST("/logout", authHandler.Logout)

	s.app.Use(authMiddleware.Middleware())
	s.app.GET("/auth/user", authHandler.User)
}
