package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ialiyldrm/restapi/app/controllers/post"
)

func PostRoutes(app *fiber.App) {
	// List All Posts
	app.Get("/posts", post.Index)

	// Show Post
	app.Get("/posts/:id", post.Show)


	// Create Post
	app.Post("/posts", post.Create)

	// Update Post
	app.Put("/post/:id",post.Update)

	// Delete Post
	app.Delete("/post/:id",post.Delete)
}