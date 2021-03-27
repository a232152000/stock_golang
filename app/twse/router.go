package twse

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine)  {
	// 分组
	v1 := e.Group("/twse")
	{
		v1.GET("/getStock/:noStr",getStockHander)
	}
}