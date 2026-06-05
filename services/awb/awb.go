package awb

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

func (s *Service) Print(ctx context.Context, req *types.PrintAWBRequest) (*types.PrintAWBResponse, error) {
	return kahttp.PostJSON[types.PrintAWBResponse](ctx, s.client, "/api/mitra/v6.1/awb/print", req)
}
