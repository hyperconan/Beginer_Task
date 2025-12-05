package orms

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	dsn := "root:123456@tcp(localhost:13306)/blog_sys?charset=utf8&parseTime=True&loc=Local"
	instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("open db failed, err:", err)
	}

	Db = instance
	Db.Migrator().AutoMigrate(&User{}, &Post{}, &Comment{})
	fmt.Println("init db success!")
}
