package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"stock/app/line"
	"stock/app/twse"
	"stock/routers"
	"stock/schedule"
)

func main() {
	//加入排程
	c := cron.New()

	c.AddFunc("@every 1m", func() {
		schedule.GetStockFunc()//更新股票資訊
	})

	c.Start()


	//加载多个APP的路由配置
	routers.Include(twse.Routers, line.Routers)
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
