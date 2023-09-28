package order

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	No    string  `json:"no"`
	Price float64 `json:"price"`
}

// 表名称
func (a *Order) Table() string {
	return "order"
}
