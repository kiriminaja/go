package address

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

func (s *Service) Provinces() (*types.ProvinceListResponse, error) {
	return kahttp.PostJSON[types.ProvinceListResponse](s.client, "/api/mitra/province", nil)
}

func (s *Service) Cities(provinsiID int) (*types.CityListResponse, error) {
	return kahttp.PostJSON[types.CityListResponse](s.client, "/api/mitra/city", map[string]any{
		"provinsi_id": provinsiID,
	})
}

func (s *Service) Districts(kabupatenID int) (*types.DistrictListResponse, error) {
	return kahttp.PostJSON[types.DistrictListResponse](s.client, "/api/mitra/kecamatan", map[string]any{
		"kabupaten_id": kabupatenID,
	})
}

func (s *Service) SubDistricts(kecamatanID int) (*types.SubDistrictListResponse, error) {
	return kahttp.PostJSON[types.SubDistrictListResponse](s.client, "/api/mitra/kelurahan", map[string]any{
		"kecamatan_id": kecamatanID,
	})
}

func (s *Service) DistrictsByName(search string) (*types.DistrictByNameResponse, error) {
	return kahttp.PostJSON[types.DistrictByNameResponse](s.client, "/api/mitra/v2/get_address_by_name", map[string]any{
		"search": search,
	})
}
