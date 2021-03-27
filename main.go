package main

import (
	"fmt"
	"stock/app/twse"
	"stock/routers"
)

func main() {
	//加载多个APP的路由配置
	routers.Include(twse.Routers)
	// 初始化路由
	r := routers.Init()

	//r.Use(middleware.LoggerToFile())

	//任何地方引入，就可以寫Log到檔案去
	//logger := logrus.New()
	//middleware.LoggerToFileSelf(logger,"animal","A walrus appears")


	if err := r.Run(":80"); err != nil {
		fmt.Println("main執行失敗, err:%v\n", err)
	}
}
