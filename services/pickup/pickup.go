package pickup

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

func (s *Service) Schedules() (*types.PickupSchedulesResponse, error) {
	return kahttp.PostJSON[types.PickupSchedulesResponse](s.client, "/api/mitra/v2/schedules", nil)
}
