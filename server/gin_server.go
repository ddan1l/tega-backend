package server

import (
	"fmt"

	"github.com/ddan1l/tega-backend/config"
	"github.com/ddan1l/tega-backend/database"
	auth_handler "github.com/ddan1l/tega-backend/handlers/auth"
	token_repository "github.com/ddan1l/tega-backend/repositories/token"
	user_repository "github.com/ddan1l/tega-backend/repositories/user"
	auth_usecase "github.com/ddan1l/tega-backend/usecases/auth"
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
	s.initializeAuthHandler()

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)

	s.app.Run(serverUrl)
}

func (s *ginServer) initializeAuthHandler() {
	userRepository := user_repository.NewUserPgRepository(s.db)
	tokenRepository := token_repository.NewTokenPgRepository(s.db)

	authUseCase := auth_usecase.NewAuthUsecaseImpl(
		userRepository,
		tokenRepository,
	)

	authHandler := auth_handler.NewAuthHandler(
		authUseCase,
	)

	s.app.POST("/register", authHandler.Register)
}
