package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"pcconf.viktorir/internal/database/postgresql"
	"pcconf.viktorir/internal/router"
)

func Run() {
	err := postgresql.Connect()
	if err != nil {
		log.Fatal("Error PostgreSQL connection! ", err)
	}
	log.Println("PostgreSQL connection!")

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost, http://localhost:8080, http://localhost:8081",
	}))

	router.Setup(app)

	log.Fatal(app.Listen(":" + "8083"))
}
