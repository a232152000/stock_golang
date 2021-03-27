package routers

import (
	"github.com/gin-gonic/gin"
	"stock/middleware"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.New()

	//每個請求將會寫入
	r.Use(middleware.LoggerToFile())

	for _, opt := range options {
		opt(r)
	}
	return r
}