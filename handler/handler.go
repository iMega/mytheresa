package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/imega/mytheresa/domain"
)

type Handler struct {
	Shop domain.Shop
}

func (handler *Handler) Products(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(`[]`))
}

func (handler *Handler) AddProduct(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(resp, "only supports POST method", http.StatusBadRequest)

		return
	}

	if !strings.Contains(req.Header.Get("Content-Type"), "application/json") {
		http.Error(resp, "only supports application/json", http.StatusBadRequest)

		return
	}

	ctx := req.Context()

	var offer Offer
	if err := json.NewDecoder(req.Body).Decode(&offer); err != nil {
		http.Error(
			resp,
			fmt.Sprintf("failed to decode body, %s", err.Error()),
			http.StatusBadRequest,
		)

		return
	}

	product := domain.Product{
		SKU:      offer.SKU,
		Name:     offer.Name,
		Category: offer.Category,
		Price: domain.Money{
			Units:    uint64(offer.Price.Original),
			Currency: offer.Price.Currency,
		},
	}

	if err := handler.Shop.Add(ctx, product); err != nil {
		http.Error(
			resp,
			fmt.Sprintf("failed to add product, %s", err.Error()),
			http.StatusBadRequest,
		)

		return
	}

	resp.Write([]byte(`ok!\n`))
}

func (handler *Handler) Healthcheck(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusNoContent)
}
