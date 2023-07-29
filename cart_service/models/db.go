package models

type CartItems struct {
	CartItemID    string `gorm:"column:cart_item_id"`
	ProductItemID string `gorm:"primaryKey,column:product_item_id"`
	Quantity      string `gorm:"column:quantity"`
	UserID        string `gorm:"primaryKey,column:user_id"`
}
