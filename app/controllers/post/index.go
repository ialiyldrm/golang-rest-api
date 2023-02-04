package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ialiyldrm/restapi/app/models"
	"github.com/ialiyldrm/restapi/platforms/database"
)

func Index(ctx *fiber.Ctx) error {
	posts := []models.Post{}
	database.Conn.Find(&posts)

	return ctx.JSON(posts)
}