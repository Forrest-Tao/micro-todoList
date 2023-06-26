package dao

import (
	"github.com/jinzhu/gorm"
	"user/db/model"
)

var DB *gorm.DB

func migrate() {
	DB.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&model.User{})
}
