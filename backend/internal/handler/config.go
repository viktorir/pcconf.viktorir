package handler

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"pcconf.viktorir/internal/model"
	"strconv"
)

func GetConfigs(ctx *fiber.Ctx) error {
	pageParam := ctx.Query("page", "1")
	limitParam := ctx.Query("limit", "15")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid page parameter")
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid limit parameter")
	}

	configs, err := model.GetConfigs(page, limit)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"configs": configs})
}

func GetConfig(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id parameter")
	}

	config, err := model.GetConfig(id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fiber.ErrNotFound
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"config": config})
}

func CreateConfig(ctx *fiber.Ctx) error {
	var config model.PCConfiguration
	if err := ctx.BodyParser(&config); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	id, err := model.CreateConfig(config)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create configuration: "+err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func UpdateConfig(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid id parameter")
	}

	var config model.PCConfiguration
	if err := ctx.BodyParser(&config); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	config.ID = id

	err = model.UpdateConfig(config)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fiber.ErrNotFound
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update configuration: "+err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}

func DeleteConfig(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID parameter")
	}

	err = model.DeleteConfig(id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fiber.ErrNotFound
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete configuration: "+err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": id,
	})
}
