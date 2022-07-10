package acceptance

import (
	"encoding/json"
	"net/http"

	"github.com/imega/mytheresa/handler"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(`
    All products from boots category
    and filtered by priceLessThan 800 euro.

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
			"http://app:8080/products?category=boots&priceLessThan=80000",
			nil,
		)
		Expect(err).NotTo(HaveOccurred())

		resp, err := http.DefaultClient.Do(req)
		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()

		var actual []handler.Offer
		err = json.NewDecoder(resp.Body).Decode(&actual)
		Expect(err).NotTo(HaveOccurred())

		expected := []handler.Offer{
			{
				SKU:      "000003",
				Name:     "Ashlington leather ankle boots",
				Category: "boots",
				Price: handler.Price{
					Currency: "EUR", Original: 71000, Final: 49700, Discount: "30%",
				},
			},
		}

		Expect(actual).To(Equal(expected))
	})
})
