package acceptance

import (
	"encoding/json"
	"net/http"

	"github.com/imega/mytheresa/handler"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(`
    All products from boots category.

    Products in the boots category have a 30% discount.
    The product with sku = 000003 has a 15% discount.
    When multiple discounts collide, the biggest discount must be applied.

    The boots category has a higher 30% discount
    than SKU 00003 which has a 15% discount
    will be applied 30%.
`, func() {
	It("products have been received", func() {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://app:8080/products?category=boots",
			nil,
		)
		Expect(err).NotTo(HaveOccurred())

		resp, err := http.DefaultClient.Do(req)
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()

		var actual []handler.Offer
		err = json.NewDecoder(resp.Body).Decode(&actual)
		Expect(err).NotTo(HaveOccurred())

		discount30 := "30%"
		expected := []handler.Offer{
			{
				SKU:      "000001",
				Name:     "BV Lean leather ankle boots",
				Category: "boots",
				Price: handler.Price{
					Original: 89000, Final: 62300, Currency: "EUR", Discount: &discount30,
				},
			},
			{
				SKU:      "000002",
				Name:     "BV Lean leather ankle boots",
				Category: "boots",
				Price: handler.Price{
					Currency: "EUR", Original: 99000, Final: 69300, Discount: &discount30,
				},
			},
			{
				SKU:      "000003",
				Name:     "Ashlington leather ankle boots",
				Category: "boots",
				Price: handler.Price{
					Currency: "EUR", Original: 71000, Final: 49700, Discount: &discount30,
				},
			},
		}

		Expect(actual).To(Equal(expected))
	})
})
