package main

import (
	"log"
	"net/http"

	"github.com/imega/mytheresa/handler"
	"github.com/imega/mytheresa/shop"
	"github.com/imega/mytheresa/storage"
)

func main() {
	handler := &handler.Handler{
		Shop: shop.New(
			storage.New(),
			shop.NewDiscounter(shop.DefaultRulesLoyaltyProgram()),
		),
	}

	http.HandleFunc("/products", handler.Products)
	http.HandleFunc("/addproduct", handler.AddProduct)
	http.HandleFunc("/healthcheck", handler.Healthcheck)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start http-server, %s", err.Error())
	}
}
