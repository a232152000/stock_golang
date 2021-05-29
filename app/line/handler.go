package line

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"os"
	"stock/config"
)

func getCallbackHander(c *gin.Context) {

	bot, err := linebot.New(
		os.Getenv(config.ViperEnvVariable("CHANNEL_SECRET")),
		os.Getenv(config.ViperEnvVariable("CHANNEL_TOKEN")),
	)
	if err != nil {
		//log.Fatal(err)
	}


	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			//log.Print(err)
		}
		return
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


//func RecommandVtuber(reqmessage string) (message string) {
//
//	url := GetApiUrl()
//	res, err := http.Get(string(url))
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer res.Body.Close()
//
//	body, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	var vtuber []VtuberInfo
//
//	if err := json.Unmarshal(body, &vtuber); err != nil {
//		log.Fatal(err)
//	}
//	rand.Seed(time.Now().UnixNano())
//	num := rand.Intn(len(vtuber))
//	if reqmessage == "hello" {
//		return "hello"
//	} else {
//		message = vtuber[num].Name + "はいいぞ!" + GetYoutubeUrl() + vtuber[num].ChannelId
//	}
//	return message
//}
