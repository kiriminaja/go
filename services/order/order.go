package order

import (
	kahttp "github.com/kiriminaja/go/http"
	"github.com/kiriminaja/go/services/order/express"
	"github.com/kiriminaja/go/services/order/instant"
)

type Service struct {
	Express *express.Service
	Instant *instant.Service
}

func New(client *kahttp.Client) *Service {
	return &Service{
		Express: express.New(client),
		Instant: instant.New(client),
	}
}
