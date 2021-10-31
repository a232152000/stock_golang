package schedule

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"stock/app/twse"
	"stock/config"
)

func GetStockFunc() {
	var stockLatest []twse.StockLatest
	var noStr string

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.ViperEnvVariable("USERNAME"), config.ViperEnvVariable("PASSWORD"), config.ViperEnvVariable("NETWORK"), config.ViperEnvVariable("SERVER"), config.ViperEnvVariable("PORT"), config.ViperEnvVariable("DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
	}

	db.Find(&stockLatest)

	for k,selectStockLatest := range stockLatest{
		if k == 0 {
			noStr = selectStockLatest.Ex + "_" + selectStockLatest.Code + ".tw"
		}else {
			noStr = noStr + ("|" + selectStockLatest.Ex + "_" + selectStockLatest.Code + ".tw")
		}
	}

	stock := twse.GetStockInfo(noStr)

	if stock == nil{
		return
	}

	//json解析成struct
	var getStockStruct twse.StockInfoStruct
	getStockUnJsonErr := json.Unmarshal(stock, &getStockStruct)
	if getStockUnJsonErr != nil {
		fmt.Println(getStockUnJsonErr)
		return
	}

	for _,v:=range getStockStruct.Msgarray {
		twse.UpdateStockLatest(v,db)
	}
}

