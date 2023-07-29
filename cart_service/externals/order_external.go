package externals

import (
	"Ecommerce/order_service/api"
	api2 "Ecommerce/product_service/api"
)

type Dependency struct {
	OrderService   api.OrderService
	ProductService api2.ProductService
}

func NewCartDependency_OrderSvc(orderService api.OrderService, productService api2.ProductService) *Dependency {
	return &Dependency{
		OrderService:   orderService,
		ProductService: productService,
	}
}
