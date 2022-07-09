package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"stock/app/line"
	"stock/app/twse"
	"stock/routers"
	"stock/schedule"
	"stock/schedule/line"
)

func main() {
	//加入排程
	c := cron.New(cron.WithSeconds())

	c.AddFunc("0 */3 8-14 * * 1-5", func() {
		schedule.GetStockFunc()//更新股票資訊
	})

	c.AddFunc("0 15 13 * * *", func() {
		lineSchedule.SendStockInformationFlexFunc() //寄送股票資訊
	})

	c.AddFunc("30 30 9 * * 1-5", func() {
		lineSchedule.SendStockInformationFlexFunc() //寄送股票資訊
	})

	c.AddFunc("30 30 11 * * 1-5", func() {
		lineSchedule.SendStockInformationFlexFunc() //寄送股票資訊
	})

	c.Start()
	defer c.Stop()

	//加载多个APP的路由配置
	routers.Include(twse.Routers, line.Routers)
	// 初始化路由
	r := routers.Init()

	//r.Use(middleware.LoggerToFile())

	//任何地方引入，就可以寫Log到檔案去
	//logger := logrus.New()
	//middleware.LoggerToFileSelf(logger,"animal","A walrus appears")


	if err := r.Run(":8011"); err != nil {
		fmt.Println("main執行失敗, err:%v\n", err)
	}
}
