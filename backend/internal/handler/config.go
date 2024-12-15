package handler

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"pcconf.viktorir/internal/model"
	"strconv"
)

func GetFullConfigs(ctx *fiber.Ctx) error {
	configs, err := model.GetFullConfigs(1, 15)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": configs})
}

func GetConfig(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		return err
	}

	config, err := model.GetConfig(id)
	if err == sql.ErrNoRows {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Config not found"})
	} else if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": config})
}
