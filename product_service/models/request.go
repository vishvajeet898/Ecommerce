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
	VariantValueIDs []string `json:"variantValueIDs"`
}

type GetAllProductItemsByVariantIDRequest struct {
	VariantId string `json:"variant_id"`
}

type UpdateProductItemRequest struct {
	ProductItemId string    `json:"productItemId"`
	Name          string    `json:"name"`
	Price         string    `json:"price"`
	Variants      []Variant `json:"variants"`
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
