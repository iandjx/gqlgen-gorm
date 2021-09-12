package dbmodel

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ProductCode string
	ProductName string
	Quantity    int
	OrderID     int
	Order       Order `gorm:"foreignKey:OrderID"`
}

type Order struct {
	gorm.Model
	CustomerName string
	OrderAmount  float64
	Items        []Item
}
