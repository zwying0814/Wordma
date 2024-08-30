package main

import (
	"wordma/cmd"
	"wordma/config"
	"wordma/log"
	"wordma/server"
	"wordma/server/model"
)

func main() {
	// 初始化日志
	log.InitLog()

	// 初始化配置
	config.InitConfigFile()
	// 连接数据库
	model.InitDatabase()

	if config.DevelopMode {
		// 开发模式则直接执行server服务
		_, err := server.Bootstrap()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// 非开发模式则执行命令行
		app := cmd.App()
		if err := app.LunchApp(); err != nil {
			log.Fatal(err)
		}
	}

}
