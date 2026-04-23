package kiriminaja

import (
	"net/http"
	"time"

	"github.com/kiriminaja/go/config"
	kahttp "github.com/kiriminaja/go/http"
	"github.com/kiriminaja/go/services/address"
	"github.com/kiriminaja/go/services/courier"
	coveragearea "github.com/kiriminaja/go/services/coverage_area"
	"github.com/kiriminaja/go/services/credit"
	"github.com/kiriminaja/go/services/order"
	"github.com/kiriminaja/go/services/payment"
	"github.com/kiriminaja/go/services/pickup"
)

type Env = config.Env
type Config = config.Config

const (
	EnvSandbox    = config.EnvSandbox
	EnvProduction = config.EnvProduction
)

const defaultTimeout = 30 * time.Second

type Client struct {
	Address      *address.Service
	Courier      *courier.Service
	CoverageArea *coveragearea.Service
	Credit       *credit.Service
	Order        *order.Service
	Payment      *payment.Service
	Pickup       *pickup.Service

	httpClient *kahttp.Client
}

func New(cfg Config) *Client {
	baseURL := cfg.BaseURL
	if baseURL == "" {
		if u, ok := config.EnvURL[cfg.Env]; ok {
			baseURL = u
		} else {
			baseURL = config.EnvURL[config.EnvSandbox]
		}
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultTimeout}
	}

	kc := &kahttp.Client{
		BaseURL:    baseURL,
		APIKey:     cfg.APIKey,
		HTTPClient: httpClient,
	}

	return &Client{
		Address:      address.New(kc),
		Courier:      courier.New(kc),
		CoverageArea: coveragearea.New(kc),
		Credit:       credit.New(kc),
		Order:        order.New(kc),
		Payment:      payment.New(kc),
		Pickup:       pickup.New(kc),
		httpClient:   kc,
	}
}
