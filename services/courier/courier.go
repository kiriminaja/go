package courier

import (
	kahttp "github.com/kiriminaja/go/http"
	"github.com/kiriminaja/go/types"
)

type Service struct {
	client *kahttp.Client
}

func New(client *kahttp.Client) *Service {
	return &Service{client: client}
}

func (s *Service) List() (*types.CourierListResponse, error) {
	return kahttp.PostJSON[types.CourierListResponse](s.client, "/api/mitra/couriers", nil)
}

func (s *Service) Group() (*types.CourierGroupResponse, error) {
	return kahttp.PostJSON[types.CourierGroupResponse](s.client, "/api/mitra/couriers_group", nil)
}

func (s *Service) Detail(courierCode string) (*types.CourierDetailResponse, error) {
	return kahttp.PostJSON[types.CourierDetailResponse](s.client, "/api/mitra/courier_services", map[string]any{
		"courier_code": courierCode,
	})
}

func (s *Service) SetWhitelistServices(services []string) (*types.SetCourierPreferenceResponse, error) {
	return kahttp.PostJSON[types.SetCourierPreferenceResponse](s.client, "/api/mitra/v3/set_whitelist_services", map[string]any{
		"services": services,
	})
}
