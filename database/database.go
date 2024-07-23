package database

import (
	"github.com/your/repo/schema"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Init() {
	var err error
	dsn := "test:1234@tcp(127.0.0.1:3307)/mysqlDB?charset=utf8mb4&parseTime=True&loc=Local"
	DBConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err := DBConn.AutoMigrate(MigrationEntities...); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal("Failed to Connect Database!")
	}
}

var MigrationEntities = []interface{}{
	&schema.ProductSchema{},
}
