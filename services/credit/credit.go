package credit

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

// Balance fetches the current KiriminAja credit balance for the authenticated
// mitra account from GET /api/mitra/v6.2/credit/balance.
func (s *Service) Balance(ctx context.Context) (*types.CreditBalanceResponse, error) {
	return kahttp.GetJSON[types.CreditBalanceResponse](ctx, s.client, "/api/mitra/v6.2/credit/balance")
}
