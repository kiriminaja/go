package types

type ProvinceListResponse struct {
	BaseResponse
	Datas []Province `json:"datas"`
}

type CityListResponse struct {
	BaseResponse
	Datas []City `json:"datas"`
}

type DistrictListResponse struct {
	BaseResponse
	Datas []District `json:"datas"`
}

type SubDistrictListResponse struct {
	BaseResponse
	Results []SubDistrict `json:"results"`
}

type DistrictByNameResponse struct {
	BaseResponse
	Data []AddressByNameResult `json:"data"`
}
