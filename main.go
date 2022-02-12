package main

import (
	"database/sql"
	"fmt"
	"project-go/config"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := connect(dsn)
	if err != nil {
		fmt.Println("Error occured: %s", err)
	}

}

func connect(url string) (*sql.DB, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	return db, nil
}
