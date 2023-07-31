package api

import (
	"Ecommerce/product_service/models"
	"Ecommerce/product_service/store"
	"fmt"
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
	CreateProduct(models.AddProductRequest) (*models.AddProductResponse, error)
	CreateProductItem(models.AddProductItemRequest) (*models.AddProductItemResponse, error)
	GetAllProduct() (*models.GetAllProductsResponse, error)
	GetAllProductItems() (*models.GetAllProductItemsResponse, error)
	GetProductItemByItemID(request models.GetProductItemByIDRequest) (*models.GetProductItemByIDResponse, error)
	GetAllProductItemsByVariantValueID(models.GetAllProductItemsByVariantValueIDRequest) (*models.GetAllProductItemsByVariantIDResponse, error)
	GetAllVariantValueByProductID(request models.GetAllVariantValueByProductIDRequest) (*models.GetAllVariantValueByProductIDResponse, error)
	UpdateProduct(models.UpdateProductRequest) (*models.UpdateProductResponse, error)
	UpdateProductItem(models.UpdateProductItemRequest) (*models.UpdateProductItemResponse, error)
}

func (productstore *ProductStore) CreateProduct(addProductRequest models.AddProductRequest) (*models.AddProductResponse, error) {
	//Create product
	product := models.Products{
		ProductID:   uuid.New().String(),
		ProductName: addProductRequest.Name,
	}
	if err := productstore.ProductStore.Create(product); err != nil {
		return nil, errUnableToAddProduct
	}

	for _, variant := range addProductRequest.Variants {

		//Create Variant
		productVariant := models.ProductVariants{
			ProductVariantID: uuid.New().String(),
			ProductID:        product.ProductID,
			VariantName:      variant.VariantName,
		}
		if err := productstore.ProductVariantStore.CreateVariant(productVariant); err != nil {
			return nil, errUnableToCreateVariant
		}

		//Create Variant Value
		for _, variantValue := range variant.VariantValues {
			productVariantValue := models.ProductVariantValues{
				ProductVariantValueID: uuid.New().String(),
				ProductVariantID:      productVariant.ProductVariantID,
				ProductVariantValue:   variantValue,
			}
			if err := productstore.ProductVariantValueStore.CreateVariantValue(productVariantValue); err != nil {
				return nil, errUnableToCreateVariantValue
			}
		}

	}

	return &models.AddProductResponse{
		ProductID: product.ProductID,
		Ok:        nil,
	}, nil
}

func (productstore *ProductStore) CreateProductItem(addProductItemRequest models.AddProductItemRequest) (*models.AddProductItemResponse, error) {
	//Create ProductItems
	productItem := models.ProductItems{
		ProductItemId: uuid.New().String(),
		ProductId:     addProductItemRequest.ProductId,
		Name:          addProductItemRequest.Name,
		Price:         addProductItemRequest.Price,
		Units:         addProductItemRequest.Units,
	}
	if err := productstore.ProductItemStore.CreateItem(productItem); err != nil {
		return nil, errUnableToAddProductItem
	}

	//Now for each variant add it in combination
	//For Optimization, we can have bulk inserts
	for _, varaintValueID := range addProductItemRequest.VariantValueIDs {
		productVariantCombination := models.ProductVariantCombinations{
			ProductItemId:         productItem.ProductItemId,
			ProductVariantValueID: varaintValueID,
		}

		if err := productstore.ProductVariantCombinationStore.CreateCombination(productVariantCombination); err != nil {
			return nil, errUnableToAddProductItem
		}
	}

	return &models.AddProductItemResponse{
		Ok: nil,
	}, nil
}

func (productstore *ProductStore) GetAllProductItems() (*models.GetAllProductItemsResponse, error) {
	dbProductItems, err := productstore.ProductItemStore.GetAllItems(store.QueryFilter{})
	if err != nil {
		return nil, errInternalServerError
	}

	getAllProductItemsResponse := models.GetAllProductItemsResponse{
		Items: dbProductItems,
	}
	return &getAllProductItemsResponse, nil
}

func (productstore *ProductStore) GetAllProductItemsByVariantValueID(request models.GetAllProductItemsByVariantValueIDRequest) (*models.GetAllProductItemsByVariantIDResponse, error) {

	querryFiler := store.QueryFilter{
		Table: "product_items",
		Rows:  "product_items.product_item_id, product_items.name, product_items.price,product_variant_combinations.product_variant_value_id",
		Join:  "inner join product_variant_combinations on product_items.product_item_id = product_variant_combinations.product_item_id",
		Where: "product_variant_combinations.product_variant_value_id = '" + request.VariantValueId + "';",
	}

	productVariants, err := productstore.ProductVariantValueStore.GetManyVariantValues(querryFiler)
	if err != nil {

		return nil, errInternalServerError
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
		return nil, errProductNotFound
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

	productVariants, err := productstore.ProductVariantStore.GetManyVariants(querryFiler)
	if err != nil {
		return nil, errInternalServerError
	}

	if len(productVariants) == 0 {
		return nil, errProductNotFound
	}

	return &models.GetAllVariantValueByProductIDResponse{
		ProductVariantValues: productVariants,
	}, nil
}

func (productstore *ProductStore) UpdateProduct(updateProductItemRequest models.UpdateProductRequest) (*models.UpdateProductResponse, error) {

	product := models.Products{
		ProductID:   updateProductItemRequest.ProductId,
		ProductName: updateProductItemRequest.Name,
	}

	//Updated the product
	err := productstore.ProductStore.UpdateOne(product)
	if err != nil {
		return nil, errUnableToUpdate
	}

	for _, variant := range updateProductItemRequest.Variants {
		//Updated the Variant Name
		productVariant := models.ProductVariants{
			ProductVariantID: variant.VariantID,
			ProductID:        updateProductItemRequest.ProductId,
			VariantName:      variant.VariantName,
		}
		if err := productstore.ProductVariantStore.UpdateOneVariant(productVariant); err != nil {
			return nil, errUnableToUpdate
		}

		for _, variantValue := range variant.VariantValues {
			//Updated the Variant Value
			variantValue := models.ProductVariantValues{
				ProductVariantID:      variant.VariantID,
				ProductVariantValueID: variantValue.VariantValueID,
				ProductVariantValue:   variantValue.VariantValue,
			}

			if err := productstore.ProductVariantValueStore.UpdateOneVariantValue(variantValue); err != nil {
				return nil, errUnableToUpdate
			}

		}

	}

	return &models.UpdateProductResponse{Ok: nil}, nil
}

func (productstore *ProductStore) UpdateProductItem(updateProductItemRequest models.UpdateProductItemRequest) (*models.UpdateProductItemResponse, error) {

	//Check if it exists
	productItem, err := productstore.ProductItemStore.GetOneItem(models.ProductItems{ProductItemId: updateProductItemRequest.ProductItemId})
	if err != nil {
		//Product does not exist
		return nil, errProductNotFound
	}

	product := models.ProductItems{
		ProductItemId: updateProductItemRequest.ProductItemId,
		Price:         updateProductItemRequest.Price,
		Units:         updateProductItemRequest.Units,
		Name:          updateProductItemRequest.Name,
		ProductId:     productItem.ProductId,
	}

	fmt.Printf("Product = \n%v", product)

	//Updated the product
	err = productstore.ProductItemStore.UpdateOneItem(product)
	if err != nil {
		return nil, errUnableToUpdate
	}

	return &models.UpdateProductItemResponse{
		Ok: nil,
	}, nil
}

func (productstore *ProductStore) GetAllProduct() (*models.GetAllProductsResponse, error) {

	//Get all the products
	productItems, err := productstore.ProductStore.GetMany(models.Products{})
	if err != nil {
		return nil, errInternalServerError
	}
	if len(productItems) == 0 {
		return nil, errProductNotFound
	}

	product := make([]models.Product, 0)
	for _, productItem := range productItems {

		//Now get all the product Variants
		querryFiler := store.QueryFilter{
			Table: "product_variant_values",
			Rows:  "product_variants.product_variant_id, product_variants.variant_name, product_variant_values.product_variant_value_id,  product_variant_values.product_variant_value",
			Join:  "inner join product_variants on product_variant_values.product_variant_id = product_variants.product_variant_id",
			Where: "product_variants.product_id = '" + productItem.ProductID + "'",
		}
		productVariants, err := productstore.ProductVariantStore.GetManyVariants(querryFiler)
		if err != nil {
			return nil, errInternalServerError
		}

		if len(productVariants) == 0 {
			return nil, errProductNotFound
		}

		//Now get all the productItems
		/*productItem := models.ProductItems{
			ProductId: productItem.ProductID,
		}*/

		productItemQueryFiler := store.QueryFilter{
			Where: "product_items.product_id = '" + productItem.ProductID + "'",
		}
		dbProductItems, err := productstore.ProductItemStore.GetAllItems(productItemQueryFiler)
		if err != nil {
			return nil, errInternalServerError
		}

		if len(dbProductItems) == 0 {
			return nil, errProductNotFound
		}

		product = append(product, models.Product{
			ProductID:    productItem.ProductID,
			Name:         productItem.ProductName,
			Variants:     productVariants,
			ProductItems: dbProductItems,
		})
	}

	return &models.GetAllProductsResponse{
		Products: product,
	}, nil
}
