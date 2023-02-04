package database

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var connection *gorm.DB
func Connection() *gorm.DB {
	once.Do(func() {
		connection = Init()
	})

	return connection
}

func Init() *gorm.DB {
	

	dsn := "root:root@tcp(localhost:3306)/restapi?charset=utf8mb4&parseTime=True&loc=Local"
	

	Connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil
	}

	return Connection
}
