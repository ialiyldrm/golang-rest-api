package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ialiyldrm/restapi/app/controllers/post"
)

func PostRoutes(app *fiber.App) {
	// List All Posts
	app.Get("/posts", post.Index)

	// Create Post
	app.Post("/post", post.Create)
}