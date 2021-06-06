package line

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
	"log"
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

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				fmt.Println(message.Text)
				//resmassage := RecommandVtuber(message.Text)
				//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(resmassage)).Do(); err != nil {
				//	log.Print(err)
				//}
			}
		}
	}

}
