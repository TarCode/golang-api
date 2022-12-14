package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tarcode/go-todo/database"
	"github.com/tarcode/go-todo/todo"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New())
	app.Use(logger.New())
	database.ConnectDB()
	defer database.DB.Close()

	api := app.Group("/api")

	todo.Register(api, database.DB)

	log.Fatal(app.Listen(":5000"))
}
