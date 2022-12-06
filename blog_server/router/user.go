package router

import (
	v1 "blog_server/api/v1"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserGroup := Router.Group("user")
	zap.S().Info("配置用户相关的url")
	{
		UserGroup.GET("list", v1.GetUserList)
	}
}
