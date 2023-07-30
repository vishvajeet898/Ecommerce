package models

type Variant struct {
	VariantName   string   `json:"variantName"`
	VariantValues []string `json:"variantValue"`
}

type AddProductRequest struct {
	JWT      string    `json:"JWT,omitempty"`
	Name     string    `json:"name"`
	Price    string    `json:"price"`
	Units    int       `json:"units"`
	Variants []Variant `json:"variants"`
}

type AddProductItemRequest struct {
	JWT             string   `json:"JWT,omitempty"`
	Name            string   `json:"name"`
	Price           string   `json:"price"`
	ProductId       string   `json:"productId"`
	Units           int      `json:"units"`
	VariantValueIDs []string `json:"variantValueIDs"`
}

type GetAllProductItemsByVariantValueIDRequest struct {
	VariantValueId string `json:"variantValueId"`
}

type VariantValuesUpdate struct {
	VariantValueID string `json:"variantValueID"`
	VariantValue   string `json:"variantValue"`
}

type VariantUpdate struct {
	VariantID     string                `json:"variantID"`
	VariantName   string                `json:"variantName"`
	VariantValues []VariantValuesUpdate `json:"variantValue"`
}
type UpdateProductRequest struct {
	JWT       string          `json:"JWT,omitempty"`
	ProductId string          `json:"productItemId"`
	Name      string          `json:"name"`
	Price     string          `json:"price"`
	Variants  []VariantUpdate `json:"variants"`
}

type UpdateProductItemRequest struct {
	JWT           string `json:"JWT,omitempty"`
	ProductItemId string `json:"productItemId"`
	Name          string `json:"name"`
	Units         int    `json:"units"`
	Price         string `json:"price"`
}

type DeleteProductRequest struct {
	ProductId string `json:"productId"`
}

type DeleteProductItemRequest struct {
	ProductItemId string `json:"productItemId"`
}

type GetProductItemByIDRequest struct {
	ProductItemId string `json:"productItemId"`
}

type GetAllVariantValueByProductIDRequest struct {
	ProductID string `json:"productID"`
}
