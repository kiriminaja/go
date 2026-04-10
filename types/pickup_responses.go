package types

type PickupScheduleItem struct {
	Clock   string `json:"clock"`
	Until   string `json:"until"`
	Expired int    `json:"expired"`
	Libur   bool   `json:"libur"`
}

type PickupSchedulesResponse struct {
	BaseResponse
	Schedules []PickupScheduleItem `json:"schedules"`
}
