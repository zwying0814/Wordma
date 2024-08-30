package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"wordma/config"
)

var (
	DB *gorm.DB
)

// InitDatabase 初始化数据库连接
func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.DatabasePath), &gorm.Config{})
	if err != nil {
		panic("数据库初始化失败！")
	}

	if config.DevelopMode {
		//先清空表
		err = DB.Migrator().DropTable(&User{}, &Comment{}, &Site{}, &Post{})
		if err != nil {
			panic("清空表失败！")
		}
	}

	// 自动迁移数据库
	err = DB.AutoMigrate(&User{}, &Comment{}, &Site{})
	if err != nil {
		panic("数据库迁移失败！")
		return
	}

	// 创建默认数据
	//CreateAdministrator()

	//fmt.Println("数据库初始化成功！")
}
