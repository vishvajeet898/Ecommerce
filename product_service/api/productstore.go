package api

import (
	"Ecommerce/product_service/models"
	"Ecommerce/product_service/store"
	"github.com/google/uuid"
)

type ProductStore struct {
	ProductStore                   store.ProductStore
	ProductItemStore               store.ProductItemsStore
	ProductVariantStore            store.VariantStore
	ProductVariantValueStore       store.VariantValueStore
	ProductVariantCombinationStore store.ProductVariantCombinationStore
}

func NewProductStoreApi(storeDependency store.Dependency) *ProductStore {
	return &ProductStore{
		ProductStore:                   storeDependency.ProductStore,
		ProductItemStore:               storeDependency.ProductItemStore,
		ProductVariantStore:            storeDependency.ProductVariantStore,
		ProductVariantValueStore:       storeDependency.ProductVariantValueStore,
		ProductVariantCombinationStore: storeDependency.ProductVariantCombinationStore,
	}
}

type ProductService interface {
	//JWT Admin
	CreateProduct(models.AddProductRequest) error
	//JWT Admin
	CreateProductItem(models.AddProductItemRequest) error
	GetAllProductItems() (*models.GetAllProductItemsResponse, error)
	GetProductItemByItemID(request models.GetProductItemByIDRequest) (*models.GetProductItemByIDResponse, error)
	GetAllProductVariantByVariantIDs(models.GetAllProductItemsByVariantIDRequest) (*models.GetAllProductItemsByVariantIDResponse, error)
	GetAllVariantValueByProductID(request models.GetAllVariantValueByProductIDRequest) (*models.GetAllVariantValueByProductIDResponse, error)
	//TODO
	//UpdateProductItem(models.UpdateProductItemRequest) error

	//DeleteProductItem(request models.DeleteProductItemRequest) error
	//GetAllProducts() (*[]models.GetAllProductItemsResponse, error)
	//DeleteProduct(request models.DeleteProductRequest) error
}

func (productstore *ProductStore) CreateProduct(addProductRequest models.AddProductRequest) error {
	//Create product
	product := models.Products{
		ProductID:   uuid.New().String(),
		ProductName: addProductRequest.Name,
	}
	if err := productstore.ProductStore.Create(product); err != nil {
		//TODO Err encode
		return err
	}

	//Create Product Item
	/*	productItem := models.ProductItems{
			ProductItemId: uuid.New().String(),
			ProductId:     product.ProductID,
			Name:          addProductRequest.Name,
			Price:         addProductRequest.Price,
			Units:         addProductRequest.Units,
		}
		if err := productstore.ProductItemStore.CreateItem(productItem); err != nil {
			//TODO Err encode
			return err
		}*/

	for _, variant := range addProductRequest.Variants {

		//Create Variant
		productVariant := models.ProductVariants{
			ProductVariantID: uuid.New().String(),
			ProductID:        product.ProductID,
			VariantName:      variant.VariantName,
		}
		if err := productstore.ProductVariantStore.CreateVariant(productVariant); err != nil {
			//TODO Err encode
			return err
		}

		for _, variantValue := range variant.VariantValues {
			productVariantValue := models.ProductVariantValues{
				ProductVariantValueID: uuid.New().String(),
				ProductVariantID:      productVariant.ProductVariantID,
				ProductVariantValue:   variantValue,
			}
			if err := productstore.ProductVariantValueStore.CreateVariantValue(productVariantValue); err != nil {
				//TODO Err encode
				return err
			}
		}
		//Create Variant Value

	}

	return nil
}

func (productstore *ProductStore) CreateProductItem(addProductItemRequest models.AddProductItemRequest) error {
	//Create ProductItems
	productItem := models.ProductItems{
		ProductItemId: uuid.New().String(),
		ProductId:     addProductItemRequest.ProductId,
		Name:          addProductItemRequest.Name,
		Price:         addProductItemRequest.Price,
	}
	if err := productstore.ProductItemStore.CreateItem(productItem); err != nil {
		//TODO Err encode
		return err
	}

	//Now for each variant add it in combination
	//For Optimization, we can have bulk inserts
	for _, varaintValueID := range addProductItemRequest.VariantValueIDs {
		productVariantCombination := models.ProductVariantCombinations{
			ProductItemId:         productItem.ProductItemId,
			ProductVariantValueID: varaintValueID,
		}

		if err := productstore.ProductVariantCombinationStore.CreateCombination(productVariantCombination); err != nil {
			//TODO Err encode
			return err
		}
	}

	return nil
}

func (productstore *ProductStore) GetAllProductItems() (*models.GetAllProductItemsResponse, error) {
	dbProductItems, err := productstore.ProductItemStore.GetAllItems()
	if err != nil {
		return nil, err
	}

	/*	productItems := make([]models.ProductItem, 0)
		for _, productItem := range dbProductItems {
			productItems = append(productItems, models.ProductItem{
				ProductID: productItem.ProductId,
				ProductItemID: productItem.ProductItemId
				Name:      productItem.Name,
				Price:     productItem.Price,
			})
		}*/

	getAllProductItemsResponse := models.GetAllProductItemsResponse{
		Items: dbProductItems,
	}
	return &getAllProductItemsResponse, nil
}

func (productstore *ProductStore) GetAllProductVariantByVariantIDs(request models.GetAllProductItemsByVariantIDRequest) (*models.GetAllProductItemsByVariantIDResponse, error) {

	querryFiler := store.QueryFilter{
		Table: "product_items",
		Rows:  "product_items.product_item_id, product_items.name, product_items.price,product_variation_combinations.product_variant_value_id",
		Join:  "inner join product_variation_combinations on product_items.product_item_id = product_variation_combinations.product_item_id",
		Where: "product_variation_combinations.product_variant_value_id = " + request.VariantId,
	}

	//TODO ErrorEncoder
	productVariants, err := productstore.ProductVariantValueStore.GetManyVariantValues(querryFiler)
	if err != nil {
		return nil, err
	}

	return &models.GetAllProductItemsByVariantIDResponse{
		ProductVariants: productVariants,
	}, nil
}

func (productstore *ProductStore) GetProductItemByItemID(getProductItemByIDRequest models.GetProductItemByIDRequest) (*models.GetProductItemByIDResponse, error) {
	productItem := models.ProductItems{
		ProductItemId: getProductItemByIDRequest.ProductItemId,
	}
	dbProductItem, err := productstore.ProductItemStore.GetOneItem(productItem)
	if err != nil {
		return nil, err
	}
	return &models.GetProductItemByIDResponse{
		ProductItem: dbProductItem,
	}, nil
}

func (productstore *ProductStore) GetAllVariantValueByProductID(getAllVariantValueByProductIDRequest models.GetAllVariantValueByProductIDRequest) (*models.GetAllVariantValueByProductIDResponse, error) {

	querryFiler := store.QueryFilter{
		Table: "product_variant_values",
		Rows:  "product_variants.product_variant_id, product_variants.variant_name, product_variant_values.product_variant_value_id,  product_variant_values.product_variant_value",
		Join:  "inner join product_variants on product_variant_values.product_variant_id = product_variants.product_variant_id",
		Where: "product_variants.product_id = '" + getAllVariantValueByProductIDRequest.ProductID + "'",
	}

	//TODO ErrorEncoder
	productVariants, err := productstore.ProductVariantStore.GetManyVariants(querryFiler)
	if err != nil {
		return nil, err
	}

	return &models.GetAllVariantValueByProductIDResponse{
		ProductVariantValues: productVariants,
	}, nil
}

/*func (productstore *ProductStore) UpdateProductItem(models.UpdateProductItemRequest) error {
	//Create ProductItems

	productItem := models.ProductItems{
		ProductItemId: uuid.New().String(),
		ProductId:     addProductItemRequest.ProductId,
		Name:          addProductItemRequest.Name,
		Price:         addProductItemRequest.Price,
	}
	if err := productstore.ProductItemStore.CreateItem(productItem); err != nil {
		//TODO Err encode
		return err
	}

	//Now for each variant add it in combination
	//For Optimization, we can have bulk inserts
	for _, varaintValueID := range addProductItemRequest.VariantValueIDs {
		productVariantCombination := models.ProductVariantCombinations{
			ProductItemId:         productItem.ProductItemId,
			ProductVariantValueID: varaintValueID,
		}

		if err := productstore.ProductVariantCombinationStore.CreateCombination(productVariantCombination); err != nil {
			//TODO Err encode
			return err
		}
	}

	return nil
}*/
