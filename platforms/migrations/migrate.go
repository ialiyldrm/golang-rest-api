package migrations

import (
	"github.com/ialiyldrm/restapi/app/models"
	"github.com/ialiyldrm/restapi/platforms/database"
)

func Migrate() {
	database.Init().AutoMigrate(&models.Post{})	
}