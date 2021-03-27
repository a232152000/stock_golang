package twse

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"reflect"
	"stock/middleware"
)

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

	fmt.Println(reflect.TypeOf(body))

	fmt.Println(body)

	//logger := logrus.New()
	//middleware.LoggerToFileSelf(logger,string(body),"success")

	//fmt.Println(string(body))
	//c.JSON(http.StatusOK, gin.H{"data":string(body)})

}