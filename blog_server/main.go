package main

import (
	"blog_server/initialize"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	port := 8028

	// 初始化logger
	initialize.InitLogger()

	// 初始化router
	Router := initialize.Routers()

	zap.S().Debugf("启动服务器， 端口: %d", port)

	if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}
}
