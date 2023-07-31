package models

type Products struct {
	ProductID   string `gorm:"primaryKey,column:product_id"`
	ProductName string `gorm:"column:product_name"`
}

type ProductItems struct {
	ProductItemId string `json:"ProductItemId" gorm:"column:product_item_id"`
	ProductId     string `json:"ProductId" gorm:"primaryKey,column:product_id"`
	Price         string `json:"Price" gorm:"column:price"`
	Name          string `json:"Name" gorm:"primaryKey,column:name"`
	Units         int    `json:"Units" gorm:"column:units"`
}

type ProductVariants struct {
	ProductVariantID string `gorm:"column:product_variant_id"`
	ProductID        string `gorm:"primaryKey,column:product_id"`
	VariantName      string `gorm:"primaryKey,column:variant_name"`
}

type ProductVariantValues struct {
	ProductVariantValueID string `gorm:"column:product_variant_value_id"`
	ProductVariantID      string `gorm:"primaryKey,column:product_variant_id"`
	ProductVariantValue   string `gorm:"primaryKey,column:product_variant_value"`
}

type ProductVariantCombinations struct {
	ProductVariantValueID string `gorm:"primaryKey,column:product_variant_value_id"`
	ProductItemId         string `gorm:"primaryKey,column:product_item_id"`
}

type Product_Variants struct {
	ProductVariantID      string `gorm:"column:product_variant_id"`
	VariantName           string `gorm:"column:variant_name"`
	ProductVariantValueID string `gorm:"column:product_variant_value_id"`
	ProductVariantValue   string `gorm:"column:product_variant_value"`
}

type Product_VariantValues struct {
	ProductItemId         string `gorm:"column:product_item_id"`
	Price                 string `gorm:"column:name"`
	Name                  string `gorm:"column:price"`
	ProductVariantValueID string `gorm:"column:product_variant_value_id"`
}

type ProductItes_VariantValues struct {
	ProductItemId string `gorm:"column:product_item_id"`
	Price         string `gorm:"column:name"`
	Name          string `gorm:"column:price"`
}
