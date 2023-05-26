package errors

import "github.com/gin-gonic/gin"

func Error500() gin.H {
	return gin.H{"code": 500, "msg": "服务器错误，请联系管理员"}
}
