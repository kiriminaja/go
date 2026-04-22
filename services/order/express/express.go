package express

import (
	"context"
	"errors"
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

func (s *Service) Track(ctx context.Context, orderID string) (*types.ExpressTrackingResponse, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}
	return kahttp.PostJSON[types.ExpressTrackingResponse](ctx, s.client, "/api/mitra/tracking", map[string]any{
		"order_id": orderID,
	})
}

// Cancel cancels an express shipment by AWB.
// The KiriminAja API requires AWB and reason as query parameters on a POST request.
func (s *Service) Cancel(ctx context.Context, awb, reason string) (*types.CancelExpressOrderResponse, error) {
	if awb == "" {
		return nil, errors.New("awb must not be empty")
	}
	return kahttp.RequestJSONTyped[types.CancelExpressOrderResponse](ctx, s.client, "/api/mitra/v3/cancel_shipment", kahttp.RequestOptions{
		Method: http.MethodPost,
		Query:  map[string]string{"awb": awb, "reason": reason},
	})
}

func (s *Service) RequestPickup(ctx context.Context, payload types.RequestPickupPayload) (*types.KAResponse, error) {
	return kahttp.PostJSON[types.KAResponse](ctx, s.client, "/api/mitra/v6.1/request_pickup", payload)
}
