package twse

import (
	"gorm.io/gorm"
)

type StockLatest struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Code string `json:"index"`
	Ex string `json:""`
	N string `json:""`
	Nf string `json:""`
	Z float64 `json:""`
	O float64 `json:""`
	H float64 `json:""`
	L float64 `json:""`
	Y float64 `json:""`
	FinalAt string `json:""`
}

func (StockLatest) TableName() string {
	return "stock_latest"
}

func CreateStockLatest(db *gorm.DB, stockLatest *StockLatest) error {
	return db.Create(stockLatest).Error
}
