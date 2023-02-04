package post

import (
	
	"github.com/gofiber/fiber/v2"
	"github.com/ialiyldrm/restapi/app/models"
	"github.com/ialiyldrm/restapi/platforms/database"
)

func Show(ctx *fiber.Ctx) error {
	post := models.Post{}

	err := database.Connection().First(&post,"id = ?",ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(post)
}