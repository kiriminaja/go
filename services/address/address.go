package address

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

func (s *Service) Provinces(ctx context.Context) (*types.ProvinceListResponse, error) {
	return kahttp.PostJSON[types.ProvinceListResponse](ctx, s.client, "/api/mitra/province", nil)
}

func (s *Service) Cities(ctx context.Context, provinsiID int) (*types.CityListResponse, error) {
	if provinsiID <= 0 {
		return nil, errors.New("provinsiID must be greater than 0")
	}
	return kahttp.PostJSON[types.CityListResponse](ctx, s.client, "/api/mitra/city", map[string]any{
		"provinsi_id": provinsiID,
	})
}

func (s *Service) Districts(ctx context.Context, kabupatenID int) (*types.DistrictListResponse, error) {
	if kabupatenID <= 0 {
		return nil, errors.New("kabupatenID must be greater than 0")
	}
	return kahttp.PostJSON[types.DistrictListResponse](ctx, s.client, "/api/mitra/kecamatan", map[string]any{
		"kabupaten_id": kabupatenID,
	})
}

func (s *Service) SubDistricts(ctx context.Context, kecamatanID int) (*types.SubDistrictListResponse, error) {
	if kecamatanID <= 0 {
		return nil, errors.New("kecamatanID must be greater than 0")
	}
	return kahttp.PostJSON[types.SubDistrictListResponse](ctx, s.client, "/api/mitra/kelurahan", map[string]any{
		"kecamatan_id": kecamatanID,
	})
}

func (s *Service) DistrictsByName(ctx context.Context, search string) (*types.DistrictByNameResponse, error) {
	if search == "" {
		return nil, errors.New("search must not be empty")
	}
	return kahttp.PostJSON[types.DistrictByNameResponse](ctx, s.client, "/api/mitra/v2/get_address_by_name", map[string]any{
		"search": search,
	})
}
