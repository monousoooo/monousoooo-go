package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewGormDB() {
	url := "root:root@tcp(localhost:3306)/monousoooo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic(err)
	}
	// 自动建表
	err = db.AutoMigrate(&SysUser{})
	DB = db
}
