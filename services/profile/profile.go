package profile

import (
	"context"

	kahttp "github.com/kiriminaja/go/http"
	"github.com/kiriminaja/go/types"
)

type Service struct {
	client *kahttp.Client
}

func New(client *kahttp.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Get(ctx context.Context) (*types.ProfileResponse, error) {
	return kahttp.GetJSON[types.ProfileResponse](ctx, s.client, "/api/mitra/v6.2/profile")
}
