package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/imega/mytheresa/domain"
)

type Handler struct {
	Shop domain.Shop
}

// nolint: funlen
func (handler *Handler) Products(resp http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(resp, "only supports GET method", http.StatusBadRequest)

		return
	}

	resp.Header().Add("Content-Type", "application/json; charset=utf-8")

	ctx := req.Context()
	queries := req.URL.Query()

	raw := queries.Get("priceLessThan")
	if raw == "" {
		raw = "0"
	}

	priceLessThan, err := strconv.ParseUint(raw, domain.Base10, domain.Bit64)
	if err != nil {
		http.Error(
			resp,
			"failed to convert priceLessThan",
			http.StatusBadRequest,
		)

		return
	}

	request := domain.Request{
		Category:      queries.Get("category"),
		PriceLessThan: priceLessThan,
	}

	offers, err := handler.Shop.Get(ctx, request)
	if err != nil {
		http.Error(
			resp,
			fmt.Sprintf("failed to get offers, %s", err.Error()),
			http.StatusInternalServerError,
		)

		return
	}

	respOffers := []Offer{}

	for _, offer := range offers {
		respOffer := Offer{
			SKU:      offer.Product.SKU,
			Name:     offer.Product.Name,
			Category: offer.Product.Category,
			Price: Price{
				Original: int(offer.Product.Price.Units),
				Final:    int(offer.Final.Units),
				Currency: offer.Product.Price.Currency,
				Discount: offer.Discount,
			},
		}

		respOffers = append(respOffers, respOffer)
	}

	if err := json.NewEncoder(resp).Encode(respOffers); err != nil {
		http.Error(
			resp,
			fmt.Sprintf("failed to encode offers, %s", err.Error()),
			http.StatusInternalServerError,
		)

		return
	}
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

	if _, err := resp.Write([]byte(`ok!\n`)); err != nil {
		http.Error(
			resp,
			fmt.Sprintf("failed to write response, %s", err.Error()),
			http.StatusInternalServerError,
		)
	}
}

func (handler *Handler) Healthcheck(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusNoContent)
}
