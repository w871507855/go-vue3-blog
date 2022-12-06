package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetUserList(c *gin.Context) {
	zap.S().Debug("获取用户列表")
}
