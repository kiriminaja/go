package types

type CourierListItem struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type CourierGroupItem struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CourierServiceItem struct {
	Name         string  `json:"name"`
	Code         string  `json:"code"`
	CutOffTime   *string `json:"cut_off_time"`
	Volumetrik   *string `json:"volumetrik"`
	Rounded      *int    `json:"rounded"`
	CourierGroup string  `json:"courier_group"`
}

type CourierListResponse struct {
	BaseResponse
	Datas []CourierListItem `json:"datas"`
}

type CourierGroupResponse struct {
	BaseResponse
	Datas []CourierGroupItem `json:"datas"`
}

type CourierDetailResponse struct {
	BaseResponse
	Datas []CourierServiceItem `json:"datas"`
}

type SetCourierPreferenceResponse struct {
	BaseResponse
}
