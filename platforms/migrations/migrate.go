package migrations

import (
	"github.com/ialiyldrm/restapi/app/models"
	"github.com/ialiyldrm/restapi/platforms/database"
)

func Migrate() {
	
	database.Connection().Set( "gorm:table_options" , "ENGINE=InnoDB" ).AutoMigrate(&models.Post{})	
}