package lineSchedule

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"stock/config"
	"time"
	"unsafe"
)

func SendStockInformationFlexFunc() {
	//logger := logrus.New()

	//bot, err := linebot.New(
	//	config.ViperEnvVariable("CHANNEL_SECRET"),
	//	config.ViperEnvVariable("CHANNEL_TOKEN"),
	//)
	//
	//if err != nil {
	//	errVal := map[string]interface{}{
	//		"error":"new linebot error",
	//	}
	//
	//	middleware.LoggerToFileSelf(logger,errVal,err.Error())
	//	log.Fatal(err)
	//}

	//撈取DB的股票資訊
	//var stockLatest []twse.StockLatest

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.ViperEnvVariable("USERNAME"), config.ViperEnvVariable("PASSWORD"), config.ViperEnvVariable("NETWORK"), config.ViperEnvVariable("SERVER"), config.ViperEnvVariable("PORT"), config.ViperEnvVariable("DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
	}

	/*
		參考-用interface存
		https://blog.csdn.net/qq_38378384/article/details/115272127
		https://gocn.vip/topics/2121
		https://kylewbanks.com/blog/query-result-to-map-in-golang
	*/
	rows, _ := db.Raw("SELECT a.token,c.* FROM users AS a JOIN user_stock_list AS b ON a.id = b.user_id JOIN stock_latest AS c ON b.stock_latest_id = c.id ORDER BY a.id").Rows()
	res := scanRows2map(rows)
	fmt.Println(res)
	//jsonString, _ := json.Marshal(res)
	//fmt.Println(jsonString)



	//jsonString := line.MakeStockInformationFlex(stockLatest)
	//
	//contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if _, err := bot.PushMessage("U55cfec471cde3a870a91c5c372258fd1", linebot.NewFlexMessage("最新股價來囉～", contents)).Do(); err != nil {
	//	fmt.Println(err)
	//}
}

func scanRows2map(rows *sql.Rows) []map[string]string {
	res := make([]map[string]string, 0) //  定义结果 map
	colTypes, _ := rows.ColumnTypes()                 // 列信息
	var rowParam = make([]interface{}, len(colTypes)) // 传入到 rows.Scan 的参数 数组
	var rowValue = make([]interface{}, len(colTypes)) // 接收数据一行列的数组

	for i, colType := range colTypes {
		rowValue[i] = reflect.New(colType.ScanType())           // 跟据数据库参数类型，创建默认值 和类型
		rowParam[i] = reflect.ValueOf(&rowValue[i]).Interface()// 跟据接收的数据的类型反射出值的地址

	}
	// 遍历
	for rows.Next() {
		rows.Scan(rowParam...) // 赋值到 rowValue 中
		record := make(map[string]string)
		for i, colType := range colTypes {

			if rowValue[i] == nil {
				record[colType.Name()] = ""
			} else {

				fmt.Printf("%T\n", rowValue[i])

				switch j := rowValue[i].(type) {
					case []uint8:
						record[colType.Name()] = Byte2Str(rowValue[i].([]byte))
					case time.Time:
						record[colType.Name()] = rowValue[i].(time.Time).Format("2006-01-02 15:04:05")
					default:
						fmt.Printf("Param #%d is unknown\n", j)
				}
			}
		}
		res = append(res, record)
	}
	return res
}

// []byte to string
func Byte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}