package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	_config "project-go/config"
	_server "project-go/server"
)

func main() {
	config := _config.LoadConfig(".")

	db := _config.InitDB(config)
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	cacheClient := _config.NewRedisClient(config)
	defer cacheClient.Close()

	server := _server.NewServer(&config, db, cacheClient)

	err := server.Run()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
