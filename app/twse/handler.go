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
	"strconv"
	"time"
)

func getStockHander(c *gin.Context) {

	noStr := c.Param("noStr")

	stock := GetStockInfo(noStr)

	//json解析成struct
	var getStockStruct StockInfoStruct
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

	for _, v := range getStockStruct.Msgarray {
		UpdateStockLatest(v, db)
	}
}

func UpdateStockLatest(v Msgarray,db *gorm.DB) {

	Z, _ := strconv.ParseFloat(v.Z, 64)
	O, _ := strconv.ParseFloat(v.O, 64)
	H, _ := strconv.ParseFloat(v.H, 64)
	L, _ := strconv.ParseFloat(v.L, 64)
	Y, _ := strconv.ParseFloat(v.Y, 64)

	stockLatest := StockLatest{
		Code:    v.C,
		Ex:      v.Ex,
		N:       v.N,
		Nf:      v.Nf,
		Z:       Z,
		O:       O,
		H:       H,
		L:       L,
		Y:       Y,
		FinalAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	db.Model(StockLatest{}).Where("code = ?", v.C).Updates(stockLatest)

	//if err := CreateStockLatest(db, stockLatest);err != nil {
	//	panic("新增 stocks 失敗，原因為 " + err.Error())
	//}
}
