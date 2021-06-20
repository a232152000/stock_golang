package line

import (
	"stock/app/twse"
	"strconv"
)

func MakeStockInformationFlex(stockLatest []twse.StockLatest) string{

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

	return jsonString
}