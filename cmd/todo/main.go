package main

import (
	"fiberTODO/cmd/database"
	"fiberTODO/internal/config"
	"fiberTODO/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"strconv"
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
	app := fiber.New(fiber.Config{
		Prefork: true, // используем предварительное форкование для увеличения производительности
	})

	app.Use(logger.New())   // Логирование запросов
	app.Use(compress.New()) // Сжатие ответов
	app.Use(recover.New())  // Восстановление после паники
	app.Use(limiter.New())  // Лимит запросов для предотвращения DDOS атак

	// Регистрация маршрутов
	routes.RegisterProductRoutes(app)

	host := cfg.Server.Host
	port := strconv.Itoa(cfg.Server.Port)
	log.Fatal(app.Listen(host + ":" + port))
}
