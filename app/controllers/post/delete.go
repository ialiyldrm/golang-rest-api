package post

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ialiyldrm/restapi/app/models"
	"github.com/ialiyldrm/restapi/platforms/database"
)

func Delete(ctx *fiber.Ctx)	error {
	post := &models.Post{}

	err := database.Connection().First(&post, "id=?",ctx.Params("id")).Error
	if err != nil {
		return err
	}

	err = database.Connection().Delete(&post).Error
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message":"Post deleted successfully",
	})
}