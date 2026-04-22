package payment

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

func (s *Service) GetPayment(ctx context.Context, paymentID string) (*types.GetPaymentResponse, error) {
	if paymentID == "" {
		return nil, errors.New("paymentID must not be empty")
	}
	return kahttp.PostJSON[types.GetPaymentResponse](ctx, s.client, "/api/mitra/v2/get_payment", map[string]any{
		"payment_id": paymentID,
	})
}
