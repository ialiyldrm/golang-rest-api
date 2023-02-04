package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ialiyldrm/restapi/pkg/routes"
	"github.com/ialiyldrm/restapi/platforms/database"
	"github.com/ialiyldrm/restapi/platforms/migrations"
)

func CreateServer(port int) {
    // Create Fiber App
    app := fiber.New()
	routes.PostRoutes(app)

	database.Init()
	migrations.Migrate()
    // Start server
    log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}