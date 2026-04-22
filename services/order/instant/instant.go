package instant

import (
	"context"
	"errors"
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

func (s *Service) Create(ctx context.Context, payload types.InstantPickupPayload) (*types.CreateInstantPickupResponse, error) {
	return kahttp.PostJSON[types.CreateInstantPickupResponse](ctx, s.client, "/api/mitra/v4/instant/pickup/request", payload)
}

func (s *Service) Track(ctx context.Context, orderID string) (*types.InstantTrackingResponse, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}
	return kahttp.GetJSON[types.InstantTrackingResponse](ctx, s.client, fmt.Sprintf("/api/mitra/v4/instant/tracking/%s", orderID))
}

func (s *Service) Cancel(ctx context.Context, orderID string) (*types.CancelInstantOrderResponse, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}
	return kahttp.DeleteJSON[types.CancelInstantOrderResponse](ctx, s.client, fmt.Sprintf("/api/mitra/v4/instant/pickup/void/%s", orderID))
}

func (s *Service) FindNewDriver(ctx context.Context, orderID string) (*types.FindNewInstantDriverResponse, error) {
	if orderID == "" {
		return nil, errors.New("orderID must not be empty")
	}
	return kahttp.PostJSON[types.FindNewInstantDriverResponse](ctx, s.client, "/api/mitra/v4/instant/pickup/find-new-driver", map[string]any{
		"order_id": orderID,
	})
}
