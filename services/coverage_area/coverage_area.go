package coveragearea

import (
	"context"

	kahttp "github.com/kiriminaja/go/http"
	"github.com/kiriminaja/go/services/address"
	"github.com/kiriminaja/go/types"
)

type Service struct {
	addr   *address.Service
	client *kahttp.Client
}

func New(client *kahttp.Client) *Service {
	return &Service{
		addr:   address.New(client),
		client: client,
	}
}

func (s *Service) PricingExpress(ctx context.Context, payload types.PricingExpressPayload) (*types.KAResponse, error) {
	return kahttp.PostJSON[types.KAResponse](ctx, s.client, "/api/mitra/v6.1/shipping_price", payload)
}

func (s *Service) PricingInstant(ctx context.Context, payload types.PricingInstantPayload) (*types.KAResponse, error) {
	return kahttp.PostJSON[types.KAResponse](ctx, s.client, "/api/mitra/v4/instant/pricing", payload)
}

func (s *Service) Provinces(ctx context.Context) (*types.ProvinceListResponse, error) {
	return s.addr.Provinces(ctx)
}

func (s *Service) Cities(ctx context.Context, provinsiID int) (*types.CityListResponse, error) {
	return s.addr.Cities(ctx, provinsiID)
}

func (s *Service) Districts(ctx context.Context, kabupatenID int) (*types.DistrictListResponse, error) {
	return s.addr.Districts(ctx, kabupatenID)
}

func (s *Service) SubDistricts(ctx context.Context, kecamatanID int) (*types.SubDistrictListResponse, error) {
	return s.addr.SubDistricts(ctx, kecamatanID)
}

func (s *Service) DistrictsByName(ctx context.Context, search string) (*types.DistrictByNameResponse, error) {
	return s.addr.DistrictsByName(ctx, search)
}
