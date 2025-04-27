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
	user_handler "github.com/ddan1l/tega-backend/handlers/user"
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
		s.app.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	s.initializeHandlers()

	s.app.Run(fmt.Sprintf(":%d", s.conf.Server.Port))
}

func (s *ginServer) initializeHandlers() {
	// Public routes
	s.initializeAuthHandler()

	// Below Protected routes
	authMiddleware := middleware.NewAuthMiddleware(
		s.factory.CreateAuthUseCase(),
	)

	s.app.Use(authMiddleware.Middleware())

	s.initializeUserHandler()
}

func (s *ginServer) initializeAuthHandler() {
	authHandler := auth_handler.NewAuthHandler(
		s.factory.CreateAuthUseCase(),
	)

	g := s.app.Group("/api/auth")

	g.POST("/register", authHandler.Register)
	g.POST("/login", authHandler.Login)
	g.POST("/logout", authHandler.Logout)

}

func (s *ginServer) initializeUserHandler() {
	userHandler := user_handler.NewUserHandler(
		s.factory.CreateUserhUseCase(),
	)

	g := s.app.Group("/api/user")

	g.GET("/", userHandler.User)
	g.GET("/projects", userHandler.UserProjects)
}
