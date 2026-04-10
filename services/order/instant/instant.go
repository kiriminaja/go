package instant

import (
	"fmt"

	kahttp "github.com/kiriminaja/go/http"
	"github.com/kiriminaja/go/types"
)

type Service struct {
	client *kahttp.Client
}

func New(client *kahttp.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Create(payload any) (*types.CreateInstantPickupResponse, error) {
	return kahttp.PostJSON[types.CreateInstantPickupResponse](s.client, "/api/mitra/v4/instant/pickup/request", payload)
}

func (s *Service) Track(orderID string) (*types.InstantTrackingResponse, error) {
	return kahttp.GetJSON[types.InstantTrackingResponse](s.client, fmt.Sprintf("/api/mitra/v4/instant/tracking/%s", orderID))
}

func (s *Service) Cancel(orderID string) (*types.CancelInstantOrderResponse, error) {
	return kahttp.DeleteJSON[types.CancelInstantOrderResponse](s.client, fmt.Sprintf("/api/mitra/v4/instant/pickup/void/%s", orderID))
}

func (s *Service) FindNewDriver(orderID string) (*types.FindNewInstantDriverResponse, error) {
	return kahttp.PostJSON[types.FindNewInstantDriverResponse](s.client, "/api/mitra/v4/instant/pickup/find-new-driver", map[string]any{
		"order_id": orderID,
	})
}
