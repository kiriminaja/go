package types

type BaseResponse struct {
	Status bool   `json:"status"`
	Method string `json:"method"`
	Text   string `json:"text"`
	Code   int    `json:"code,omitempty"`
}