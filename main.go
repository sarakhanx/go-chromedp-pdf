package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/sarakhanx/go-try-something/config/dotenv"
	"github.com/sarakhanx/go-try-something/routes/helper"
)

func main() {
	dotenv.Init()
	app := fiber.New()

	log.Println("Server is starting...")
	route_helper.Hifolks(app)

	log.Fatal(app.Listen(":8080"))
	log.Fatal(app.Listen(":8080"))
}
