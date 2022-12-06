package initialize

import "go.uber.org/zap"

func InitLogger() {
	log, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(log)
}
