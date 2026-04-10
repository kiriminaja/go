package types

type BaseResponse struct {
	Status bool   `json:"status"`
	Method string `json:"method"`
	Text   string `json:"text"`
	Code   any    `json:"code,omitempty"`

}