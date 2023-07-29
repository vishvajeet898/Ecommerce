package models

type AddProductResponse struct {
	Ok error `json:"ok"`
}

type AddProductItemResponse struct {
	Ok error `json:"ok"`
}

type GetAllProductItemsResponse struct {
	Items []ProductItems `json:"items"`
}

type ItemVariant struct {
	ProductItemID         string `json:"ProductItemID"`
	Price                 string `json:"price"`
	Name                  string `json:"name"`
	ProductVariantValueID string `json:"productVariantValueID"`
}

type GetAllProductItemsByVariantIDResponse struct {
	ProductVariants []Product_VariantValues `json:"productVariants"`
}

/*type ProductItem struct {
	Name          string        `json:"name"`
	ProductItemID string        `json:"ProductItemID"`
	Price         string        `json:"price"`
	Variants      []ItemVariant `json:"variants"`
}
*/

type ProductItem struct {
	ProductItemID string `json:"ProductItemID"`
	ProductID     string `json:"productID"`
	Price         string `json:"price"`
	Name          string `json:"name"`
}

type GetProductItemByIDResponse struct {
	ProductItem ProductItems `json:"productItem"`
}

type GetAllVariantValueByProductIDResponse struct {
	ProductVariantValues []Product_Variants `json:"productVariantValues"`
}
