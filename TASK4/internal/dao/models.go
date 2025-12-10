package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hyperconan.com/blog_sys/internal/config"
	"hyperconan.com/blog_sys/internal/pkg/logger"
)

var Db *gorm.DB

func init() {
	logger.Init()
	// 加载配置文件
	cfg, err := config.LoadConfig("")
	if err != nil {
		logger.S.Warnf("加载配置文件失败，使用默认配置: %v", err)
		// 如果配置文件加载失败，使用默认配置
		dsn := "root:123456@tcp(localhost:13306)/blog_sys?charset=utf8&parseTime=True&loc=Local"
		instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.S.Fatalw("open db failed", "err", err)
			return
		}
		Db = instance
	} else {
		// 使用配置文件中的配置
		dsn := cfg.Database.GetDSN()
		instance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.S.Fatalw("open db failed with config", "err", err)
			return
		}
		Db = instance
		logger.S.Infow("使用配置文件连接数据库成功")
	}

	if err := Db.Migrator().AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		logger.S.Fatalw("auto migrate failed", "err", err)
	}
	logger.S.Infow("init db success")
}
