package line

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"stock/app/twse"
	"stock/config"
	"stock/middleware"
)

func getCallbackHander(c *gin.Context) {
	logger := logrus.New()

	bot, err := linebot.New(
		config.ViperEnvVariable("CHANNEL_SECRET"),
		config.ViperEnvVariable("CHANNEL_TOKEN"),
	)

	if err != nil {
		errVal := map[string]interface{}{
			"error":"new linebot error",
		}

		middleware.LoggerToFileSelf(logger,errVal,err.Error())
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			errVal := map[string]interface{}{
				"error":"linebot ParseRequest error",
			}

			middleware.LoggerToFileSelf(logger,errVal,err.Error())
			log.Fatal(err)
		}
	}

	//撈取DB的股票資訊
	var stockLatest []twse.StockLatest

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.ViperEnvVariable("USERNAME"), config.ViperEnvVariable("PASSWORD"), config.ViperEnvVariable("NETWORK"), config.ViperEnvVariable("SERVER"), config.ViperEnvVariable("PORT"), config.ViperEnvVariable("DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
	}

	db.Find(&stockLatest)

	jsonString := MakeStockInformationFlex(stockLatest)

	for _, event := range events {
		//userID := event.Source.UserID
		replyToken := event.ReplyToken
		//groupID := event.Source.GroupID
		//RoomID := event.Source.RoomID

		if event.Type == linebot.EventTypeMessage {
			//switch message := event.Message.(type) {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//fmt.Println(userID)
				fmt.Println(message.Text)

				contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))

				if err != nil {
					fmt.Println(err)
				}
				if _, err := bot.ReplyMessage(replyToken, linebot.NewFlexMessage("最新股價來囉～", contents)).Do(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}

}
