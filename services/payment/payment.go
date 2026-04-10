package payment

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

func (s *Service) GetPayment(paymentID string) (*types.GetPaymentResponse, error) {
	return kahttp.PostJSON[types.GetPaymentResponse](s.client, "/api/mitra/v2/get_payment", map[string]any{
		"payment_id": paymentID,
	})
}
