package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
	"pcconf.viktorir/internal/handler"
)

func Setup(app *fiber.App) {
	api := app.Group("/api", logger.New(logger.Config{
		Format:     "{\"time\":\"${time}\",\"method\":\"${method}\",\"path\":\"${path}\",\"status\":${status},\"user-agent\":\"${ua}\",\"route\":\"${route}\",\"error\":\"${error}\"}\n",
		TimeFormat: "02/Jan/2006:15:04:05",
		TimeZone:   "Local",
		Output:     os.Stdout,
	}))

	v1 := api.Group("/v1")
	v1.Get("/ping", handler.Ping)

	configs := v1.Group("/configs")
	configs.Get("/", handler.GetConfigs)
	configs.Get("/:id", handler.GetConfig)
	configs.Post("/", handler.CreateConfig)
	configs.Put("/:id", handler.UpdateConfig)
	configs.Delete("/:id", handler.DeleteConfig)

}
