package express

import (
	"net/http"

	kahttp "github.com/kiriminaja/go/http"
	"github.com/kiriminaja/go/types"
)

type Service struct {
	client *kahttp.Client
}

func New(client *kahttp.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Track(orderID string) (*types.ExpressTrackingResponse, error) {
	return kahttp.PostJSON[types.ExpressTrackingResponse](s.client, "/api/mitra/tracking", map[string]any{
		"order_id": orderID,
	})
}

func (s *Service) Cancel(awb, reason string) (*types.CancelExpressOrderResponse, error) {
	return kahttp.RequestJSONTyped[types.CancelExpressOrderResponse](s.client, "/api/mitra/v3/cancel_shipment", kahttp.RequestOptions{
		Method: http.MethodPost,
		Query:  map[string]string{"awb": awb, "reason": reason},
	})
}

func (s *Service) RequestPickupV5(payload any) (*types.KAResponse, error) {
	return kahttp.PostJSON[types.KAResponse](s.client, "/api/mitra/v5/request_pickup", payload)
}

func (s *Service) RequestPickupV61(payload any) (*types.KAResponse, error) {
	return kahttp.PostJSON[types.KAResponse](s.client, "/api/mitra/v6.1/request_pickup", payload)
}
