package line

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine)  {
	// 分组
	v1 := e.Group("/line")
	{
		v1.GET("/callback",getCallbackHander)
	}
}