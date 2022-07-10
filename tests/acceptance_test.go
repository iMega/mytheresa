package acceptance

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/imega/mytheresa/handler"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	err := WaitForSystemUnderTestReady()
	Expect(err).NotTo(HaveOccurred())

	err = createStore()
	Expect(err).NotTo(HaveOccurred())
})

func WaitForSystemUnderTestReady() error {
	req, err := http.NewRequest(http.MethodGet, "http://app:8080/healthcheck", nil)
	if err != nil {
		return err
	}

	for attempts := 10; attempts > 0; attempts-- {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		if err == nil && resp != nil && resp.StatusCode == http.StatusNoContent {
			return nil
		}

		log.Printf("ATTEMPTING TO CONNECT")

		<-time.After(time.Second)
	}

	return errors.New("SUT is not ready for tests")
}

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
}

func createStore() error {
	offers := []handler.Offer{
		{
			SKU:      "000001",
			Name:     "BV Lean leather ankle boots",
			Category: "boots",
			Price:    handler.Price{Original: 89000, Currency: "EUR"},
		},
		{
			SKU:      "000002",
			Name:     "BV Lean leather ankle boots",
			Category: "boots",
			Price:    handler.Price{Currency: "EUR", Original: 99000},
		},
		{
			SKU:      "000003",
			Name:     "Ashlington leather ankle boots",
			Category: "boots",
			Price:    handler.Price{Currency: "EUR", Original: 71000},
		},
		{
			SKU:      "000004",
			Name:     "Naima embellished suede sandals",
			Category: "sandals",
			Price:    handler.Price{Currency: "EUR", Original: 79500},
		},
		{
			SKU:      "000005",
			Name:     "Nathane leather sneakers",
			Category: "sneakers",
			Price:    handler.Price{Currency: "EUR", Original: 59000},
		},
		{
			SKU:      "000006",
			Name:     "product over limit",
			Category: "limit",
			Price:    handler.Price{Currency: "EUR", Original: 99999999},
		},
	}

	for _, offer := range offers {
		buf := bytes.NewBuffer(nil)

		if err := json.NewEncoder(buf).Encode(offer); err != nil {
			return fmt.Errorf("failed to encode offer, %w", err)
		}

		http.DefaultClient.Post("http://app:8080/addproduct", "application/json", buf)
	}

	return nil
}
