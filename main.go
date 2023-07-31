package main

import (
	api5 "Ecommerce/address_service/api"
	store5 "Ecommerce/address_service/store"
	api4 "Ecommerce/cart_service/api"
	"Ecommerce/cart_service/externals"
	store4 "Ecommerce/cart_service/store"
	"Ecommerce/middleware/jwt"
	api3 "Ecommerce/order_service/api"
	store3 "Ecommerce/order_service/store"
	api2 "Ecommerce/product_service/api"
	store2 "Ecommerce/product_service/store"
	"Ecommerce/user_service/api"
	"Ecommerce/user_service/store"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
)

type configuration struct {
	ServerAddress string
	DBAddress     string
}

func main() {
	var config configuration
	err := envconfig.Process("MYAPP", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	token, err := jwt.NewToken("admin_1", jwt.AdminScope)
	println(token)
	println()

	//dbUrl := "postgres://vishwajeet:ecommerce@localhost:5434/EcommerceDB-New-New"
	dbUrl := "postgres://vishwajeet:PpNpjrCAerMjyclJWNOsOe8dcBVuthPW@dpg-cj1982c07spjv9rabnp0-a.oregon-postgres.render.com/productdb_32q0"
	db, _ := gorm.Open(postgres.Open( /*config.DBAddress*/ dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	//AutoMigrate

	//User Api entity
	userEntity := api.NewUserStoreApi(store.Dependency{
		UsersStore: store.NewEntityStore(db),
	})
	userEndpoints := api.MakeUserEndpoints(userEntity)

	//Product Api entity
	productEntity := api2.NewProductStoreApi(store2.Dependency{
		ProductStore:                   store2.NewEntityStore(db),
		ProductItemStore:               store2.NewEntityStore(db),
		ProductVariantValueStore:       store2.NewEntityStore(db),
		ProductVariantStore:            store2.NewEntityStore(db),
		ProductVariantCombinationStore: store2.NewEntityStore(db),
	})
	productEndpoints := api2.MakeProductEndpoints(productEntity)

	//Order Api entity
	orderEntity := api3.NewOrderStoreApi(store3.Dependency{
		OrderItemStore: store3.NewEntityStore(db),
		OrderStore:     store3.NewEntityStore(db),
	})
	orderEndpoints := api3.MakeOrderEndpoints(orderEntity)

	//Cart Api entity
	cartEntity := api4.NewCartStoreApi(store4.Dependency{
		CartItemStore: store4.NewEntityStore(db),
	}, externals.Dependency{
		OrderService:   orderEntity,
		ProductService: productEntity,
	})
	cartEndPoints := api4.MakeOrderEndpoints(cartEntity)

	//Address Api entity
	addressEntity := api5.NewAddressStoreApi(store5.Dependency{
		AddressStore: store5.NewEntityStore(db),
	})
	addressEndpoints := api5.MakeAddressEndpoints(addressEntity)

	mux := mux2.NewRouter()
	SubRouter := mux.PathPrefix("/").Subrouter()
	mux.Handle("/", api.NewHttpService(userEndpoints, SubRouter))

	mux.Handle("/", api2.NewHttpService(productEndpoints, SubRouter))

	mux.Handle("/", api3.NewHttpService(orderEndpoints, SubRouter))

	mux.Handle("/", api4.NewHttpService(cartEndPoints, SubRouter))

	mux.Handle("/", api5.NewHttpService(addressEndpoints, SubRouter))

	fmt.Printf("listening on 7171")
	http.ListenAndServe(":7171", mux)

}
