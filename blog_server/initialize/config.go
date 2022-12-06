package initialize

import (
	"blog_server/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	path, _ := os.Getwd()
	fmt.Println("当前路径", path)
	debug := GetEnvInfo("BLOG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-pro", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("%s-debug", configFilePrefix)
	}
	v := viper.New()
	// 配置文件路径
	v.AddConfigPath(path + "/config")
	v.SetConfigName(configFileName)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	err := v.Unmarshal(global.ServerConfig)
	if err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息：&v", global.ServerConfig)

	// viper动态监控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化： %s", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.ServerConfig)
		zap.S().Infof("配置信息：&v", global.ServerConfig)
	})
}
