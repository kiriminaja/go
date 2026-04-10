package coveragearea

import (
	kahttp "github.com/kiriminaja/go/http"
	"github.com/kiriminaja/go/services/address"
	"github.com/kiriminaja/go/types"
)

type Service struct {
	*address.Service
	client *kahttp.Client
}

func New(client *kahttp.Client) *Service {
	return &Service{
		Service: address.New(client),
		client:  client,
	}
}

func (s *Service) PricingExpress(payload types.PricingExpressPayload) (*types.KAResponse, error) {
	return kahttp.PostJSON[types.KAResponse](s.client, "/api/mitra/v6.1/shipping_price", payload)
}

func (s *Service) PricingInstant(payload types.PricingInstantPayload) (*types.KAResponse, error) {
	return kahttp.PostJSON[types.KAResponse](s.client, "/api/mitra/v4/instant/pricing", payload)
}
