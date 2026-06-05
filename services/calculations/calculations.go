package calculations

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

func (s *Service) COD(ctx context.Context, req *types.CalculateCODRequest) (*types.CalculateCODResponse, error) {
	return kahttp.PostJSON[types.CalculateCODResponse](ctx, s.client, "/api/mitra/calculations/cod", req)
}
