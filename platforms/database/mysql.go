package database

import (

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Init() error {
	var err error

	dsn := "root:root@tcp(localhost:3306)/restapi?charset=utf8mb4&parseTime=True&loc=Local"
	

	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	return nil
}
