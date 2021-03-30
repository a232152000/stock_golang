package twse

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	//"stock/config"
	"stock/middleware"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
)

type GetStockInfo struct {
	Msgarray []struct {
		Tv    string `json:"tv"`
		Ps    string `json:"ps"`
		Pz    string `json:"pz"`
		Fv    string `json:"fv"`
		Oa    string `json:"oa"`
		Ob    string `json:"ob"`
		A     string `json:"a"`
		B     string `json:"b"`
		C     string `json:"c"`
		D     string `json:"d"`
		Ch    string `json:"ch"`
		Ot    string `json:"ot"`
		Tlong string `json:"tlong"`
		F     string `json:"f"`
		IP    string `json:"ip"`
		G     string `json:"g"`
		Mt    string `json:"mt"`
		Ov    string `json:"ov"`
		H     string `json:"h"`
		I     string `json:"i"`
		It    string `json:"it"`
		Oz    string `json:"oz"`
		L     string `json:"l"`
		N     string `json:"n"`
		O     string `json:"o"`
		P     string `json:"p"`
		Ex    string `json:"ex"`
		S     string `json:"s"`
		T     string `json:"t"`
		U     string `json:"u"`
		V     string `json:"v"`
		W     string `json:"w"`
		Nf    string `json:"nf"`
		Y     string `json:"y"`
		Z     string `json:"z"`
		Ts    string `json:"ts"`
	} `json:"msgArray"`
	Referer   string `json:"referer"`
	Userdelay int    `json:"userDelay"`
	Rtcode    string `json:"rtcode"`
	Querytime struct {
		Sysdate           string `json:"sysDate"`
		Stockinfoitem     int    `json:"stockInfoItem"`
		Stockinfo         int    `json:"stockInfo"`
		Sessionstr        string `json:"sessionStr"`
		Systime           string `json:"sysTime"`
		Showchart         bool   `json:"showChart"`
		Sessionfromtime   int64  `json:"sessionFromTime"`
		Sessionlatesttime int64  `json:"sessionLatestTime"`
	} `json:"queryTime"`
	Rtmessage   string `json:"rtmessage"`
	Exkey       string `json:"exKey"`
	Cachedalive int    `json:"cachedAlive"`
}

func getStockHander(c *gin.Context) {

	noStr := c.Param("noStr")

	url := "https://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch="+noStr+"&json=1&delay=0"
	method := "GET"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	//json解析成struct
	var getStockStruct GetStockInfo
	getStockUnJsonErr := json.Unmarshal(body, &getStockStruct)
	if getStockUnJsonErr != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", getStockStruct.Userdelay)

	//json解析成map
	var getStockMap map[string]interface{}
	getStockUnMapErr := json.Unmarshal(body, &getStockMap)
	if getStockUnMapErr != nil {
		fmt.Println(err)
		return
	}

	logger := logrus.New()
	middleware.LoggerToFileSelf(logger,getStockMap,"success")

	//c.JSON(http.StatusOK, gin.H{"data":string(body)})

}