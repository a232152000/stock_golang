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
	"strconv"
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

/*
	{
	  "type": "carousel",
	  "contents": [
		{
		  "type": "bubble",
		  "body": {
			"type": "box",
			"layout": "vertical",
			"contents": [
			  {
				"type": "text",
				"text": "0050-元大台灣50",
				"weight": "bold",
				"size": "xl"
			  },
			  {
				"type": "box",
				"layout": "vertical",
				"margin": "lg",
				"spacing": "sm",
				"contents": [
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "當前成交價：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "138.1",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "昨收：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "137.6",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "開盤：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "138.1",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "最高：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "138.5",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "最低：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "137.85",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "漲跌：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "+0.36%",
						"wrap": true,
						"color": "#ff0000",
						"size": "sm",
						"flex": 4
					  }
					]
				  }
				]
			  }
			]
		  }
		},
		{
		  "type": "bubble",
		  "body": {
			"type": "box",
			"layout": "vertical",
			"contents": [
			  {
				"type": "text",
				"text": "00881-國泰台灣5G+",
				"weight": "bold",
				"size": "xl"
			  },
			  {
				"type": "box",
				"layout": "vertical",
				"margin": "lg",
				"spacing": "sm",
				"contents": [
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "當前成交價：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "17.65",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "昨收：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "17.63",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "開盤：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "17.69",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "最高：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "17.77",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "最低：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "17.65",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "漲跌：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "+0.11%",
						"wrap": true,
						"color": "#ff0000",
						"size": "sm",
						"flex": 4
					  }
					]
				  }
				]
			  }
			]
		  }
		}
	  ]
	}
*/

	var jsonString = `{
	"type": "carousel",
	"contents": [`

	for k,selectStockLatest := range stockLatest{
		//計算漲跌幅百分比
		priceFluctuationLimit := (selectStockLatest.Z - selectStockLatest.Y) / selectStockLatest.Y * 100

		if k > 0 {
			jsonString = jsonString + `,`
		}

		jsonString = jsonString + `{
		  "type": "bubble",
		  "body": {
			"type": "box",
			"layout": "vertical",
			"contents": [
			  {
				"type": "text",
				"text": "` + selectStockLatest.Code + `-` + selectStockLatest.N + `",
				"weight": "bold",
				"size": "xl"
			  },
			  {
				"type": "box",
				"layout": "vertical",
				"margin": "lg",
				"spacing": "none",
				"contents": [
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "當前成交價：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "` + strconv.FormatFloat(selectStockLatest.Z, 'f', 2, 64) + `",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "昨收：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "` + strconv.FormatFloat(selectStockLatest.Y, 'f', 2, 64) + `",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "開盤：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "` + strconv.FormatFloat(selectStockLatest.O, 'f', 2, 64) + `",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "最高：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "` + strconv.FormatFloat(selectStockLatest.H, 'f', 2, 64) + `",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "最低：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },
					  {
						"type": "text",
						"text": "` + strconv.FormatFloat(selectStockLatest.L, 'f', 2, 64) + `",
						"wrap": true,
						"color": "#666666",
						"size": "sm",
						"flex": 4
					  }
					]
				  },
				  {
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					  {
						"type": "text",
						"text": "漲跌：",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 2
					  },`

		//當上漲或打平為紅字，其餘為綠字
		if(priceFluctuationLimit >= 0){
			jsonString = jsonString + `{
						"type": "text",
						"text": "` + strconv.FormatFloat(priceFluctuationLimit, 'f', 2, 64) + `%",
						"wrap": true,
						"color": "#ff0000",
						"size": "sm",
						"flex": 4
					  }`
		}else{
			jsonString = jsonString + `{
						"type": "text",
						"text": "` + strconv.FormatFloat(priceFluctuationLimit, 'f', 2, 64) + `%",
						"wrap": true,
						"color": "#00ff00",
						"size": "sm",
						"flex": 4
					  }`
		}

		jsonString = jsonString + `
					]
				  }
				]
			  }
			]
		  }
		}`
	}

	jsonString += `]}`

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
				if _, err := bot.ReplyMessage(replyToken, linebot.NewFlexMessage("Sorry :(, please update your app.", contents)).Do(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}

}
