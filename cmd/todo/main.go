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

	databaseURL := cfg.GetDatabaseURL()

	var db database.DB
	switch cfg.Database.Type {
	case "postgres":
		db = &database.Postgres{}
	default:
		log.Fatalf("Unsupported database type: %s.", cfg.Database.Type)
	}

	if err := db.InitDB(databaseURL); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.CloseDB()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(limiter.New())

	routes.RegisterProductRoutes(app, db)

	host := cfg.Server.Host
	port := strconv.Itoa(cfg.Server.Port)
	log.Fatal(app.Listen(host + ":" + port))
}
