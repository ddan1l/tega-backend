package main

import (
	"github.com/ddan1l/tega-backend/config"
	"github.com/ddan1l/tega-backend/database"
	"github.com/ddan1l/tega-backend/server"
)

func main() {
	conf := config.GetConfig()

	db := database.NewPostgresDatabase(conf)

	server.NewGinServer(conf, db).Start()
}
