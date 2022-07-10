package main

import (
	"log"
	"net/http"

	"github.com/imega/mytheresa/domain"
	"github.com/imega/mytheresa/shop"
	"github.com/imega/mytheresa/storage"
)

func main() {
	handler := &Handler{
		Shop: shop.New(
			storage.New(),
			shop.NewDiscounter(shop.DefaultRulesLoyaltyProgram()),
		),
	}

	http.HandleFunc("/products", handler.Products)
	http.HandleFunc("/addproduct", handler.Products)
	http.HandleFunc("/healthcheck", handler.Healthcheck)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start http-server, %s", err.Error())
	}
}

type Handler struct {
	Shop domain.Shop
}

func (handler *Handler) Products(resp http.ResponseWriter, req *http.Request) {
	//
}

func (handler *Handler) AddProduct(resp http.ResponseWriter, req *http.Request) {
	//
}

func (handler *Handler) Healthcheck(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(`ok!`))
}
