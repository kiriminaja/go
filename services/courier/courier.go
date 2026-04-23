package courier

import (
	"context"
	"errors"

	kahttp "github.com/kiriminaja/go/http"
	"github.com/kiriminaja/go/types"
)

type Service struct {
	client *kahttp.Client
}

func New(client *kahttp.Client) *Service {
	return &Service{client: client}
}

func (s *Service) List(ctx context.Context) (*types.CourierListResponse, error) {
	return kahttp.PostJSON[types.CourierListResponse](ctx, s.client, "/api/mitra/couriers", nil)
}

func (s *Service) Group(ctx context.Context) (*types.CourierGroupResponse, error) {
	return kahttp.PostJSON[types.CourierGroupResponse](ctx, s.client, "/api/mitra/couriers_group", nil)
}

func (s *Service) Detail(ctx context.Context, courierCode string) (*types.CourierDetailResponse, error) {
	if courierCode == "" {
		return nil, errors.New("courierCode must not be empty")
	}
	return kahttp.PostJSON[types.CourierDetailResponse](ctx, s.client, "/api/mitra/courier_services", map[string]any{
		"courier_code": courierCode,
	})
}

func (s *Service) SetWhitelistServices(ctx context.Context, services []string) (*types.SetCourierPreferenceResponse, error) {
	if len(services) == 0 {
		return nil, errors.New("services must not be empty")
	}
	return kahttp.PostJSON[types.SetCourierPreferenceResponse](ctx, s.client, "/api/mitra/v3/set_whitelist_services", map[string]any{
		"services": services,
	})
}
