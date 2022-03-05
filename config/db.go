package config

import (
	"database/sql"
	"fmt"
	"project-go/exception"
)

func InitDB(config MainConfig) *sql.DB {
	dbConfig := config.Database

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.DBUser,
		dbConfig.DBPass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	db, err := sql.Open(`mysql`, connection)
	if err != nil {
		fmt.Println("Error occured: ", err)
		exception.PanicIfNeeded(err)
	}

	return db
}
