package twse

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"stock/config"
	"stock/middleware"
)

func getStockHander(c *gin.Context) {

	noStr := c.Param("noStr")

	stock := getStockInfo(noStr)

	//json解析成struct
	var getStockStruct GetStockInfo
	getStockUnJsonErr := json.Unmarshal(stock, &getStockStruct)
	if getStockUnJsonErr != nil {
		fmt.Println(getStockUnJsonErr)
		return
	}

	//fmt.Printf("%+v", getStockStruct.Userdelay)

	//json解析成map
	var getStockMap map[string]interface{}
	getStockUnMapErr := json.Unmarshal(stock, &getStockMap)
	if getStockUnMapErr != nil {
		fmt.Println(getStockUnMapErr)
		return
	}

	logger := logrus.New()
	middleware.LoggerToFileSelf(logger,getStockMap,"success")

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.ViperEnvVariable("USERNAME"), config.ViperEnvVariable("PASSWORD"), config.ViperEnvVariable("NETWORK"), config.ViperEnvVariable("SERVER"), config.ViperEnvVariable("PORT"), config.ViperEnvVariable("DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
	}

	stockLatest := &StockLatest{
		N: "abcdef",
	}

	//if err := CreateStockLatest(db, stockLatest);err != nil {
	//	panic("新增 stock_latest 失敗，原因為 " + err.Error())
	//}

	db.Model(StockLatest{}).Where("code = ?", "aaa").Updates(stockLatest)

}