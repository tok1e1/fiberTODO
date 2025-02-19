package main

import (
	"fiberTODO/cmd/database"
	"fiberTODO/internal/config"
	"log"
)

func main() {

	cfg := config.MustLoad()

	//Реализация БД
	var db database.DB

	dbType := cfg.Database.Type
	if dbType == "" {
		log.Fatal("DB_TYPE is not set in the config. Please specify the database type.")
	}

	switch dbType {
	case "postgres":
		db = &database.Postgres{}
	default:
		log.Fatalf("Unsupported database type: %s.", dbType)
	}

	db.InitDB()
	defer db.CloseDB()

	//Реализация обработки на стороне сервера
}
