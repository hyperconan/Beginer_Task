package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hyperconan.com/blog_sys/internal/config"
)

var Db *gorm.DB

func init() {
	// 加载配置文件
	cfg, err := config.LoadConfig("")
	if err != nil {
		fmt.Printf("加载配置文件失败: %v，使用默认配置\n", err)
		// 如果配置文件加载失败，使用默认配置
		dsn := "root:123456@tcp(localhost:13306)/blog_sys?charset=utf8&parseTime=True&loc=Local"
		instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("open db failed, err:", err)
			return
		}
		Db = instance
	} else {
		// 使用配置文件中的配置
		dsn := cfg.Database.GetDSN()
		instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("open db failed, err:", err)
			return
		}
		Db = instance
		fmt.Println("使用配置文件连接数据库成功")
	}

	Db.Migrator().AutoMigrate(&User{}, &Post{}, &Comment{})
	fmt.Println("init db success!")
}
