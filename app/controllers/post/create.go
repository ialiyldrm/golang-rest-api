package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ialiyldrm/restapi/app/models"
	"github.com/ialiyldrm/restapi/platforms/database"
)

func Create(ctx *fiber.Ctx) error {
	post := models.Post{}

	if err := ctx.BodyParser(&post); err != nil {
        return ctx.Status(503).JSON(err)
    }

    if err := database.Conn.Create(&post).Error; err != nil {
        return ctx.Status(503).JSON(err)
    }

    return ctx.JSON(post)
}