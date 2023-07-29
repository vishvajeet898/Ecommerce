package models

import "time"

type Orders struct {
	OrderID   string    `gorm:"column:order_id"`
	UserID    string    `gorm:"column:user_id"`
	Address   string    `gorm:"column:address"`
	Status    string    `gorm:"column:status"`
	Timestamp time.Time `gorm:"column:timestamp"`
}

type OrderItems struct {
	OrderID       string `gorm:"column:order_id"`
	OrderItemID   string `gorm:"column:order_item_id"`
	ProductItemID string `gorm:"column:product_item_id"`
	Quantity      string `gorm:"column:quantity"`
	Price         string `gorm:"column:price"`
}
