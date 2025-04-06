package main

import (
	"github.com/ddan1l/tega-api/config"
	"github.com/ddan1l/tega-api/database"
	"github.com/ddan1l/tega-api/server"
)

func main() {
	conf := config.GetConfig()

	db := database.NewPostgresDatabase(conf)

	server.NewGinServer(conf, db).Start()
}
