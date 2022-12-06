package main

import (
	"blog_server/global"
	"blog_server/initialize"
	"fmt"
	"go.uber.org/zap"
)

func main() {

	// 初始化logger
	initialize.InitLogger()

	// 初始化配置文件
	initialize.InitConfig()
	// 初始化router
	Router := initialize.Routers()

	zap.S().Debugf("启动服务器， 端口: %d", global.ServerConfig.Port)

	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}
}
