package types

import "encoding/json"

type KAResponse struct {
	Status     bool            `json:"status"`
	Text       string          `json:"text"`
	Method     string          `json:"method,omitempty"`
	StatusCode *int            `json:"status_code,omitempty"`
	Code       *int            `json:"code,omitempty"`
	Datas      json.RawMessage `json:"datas,omitempty"`
	Data       json.RawMessage `json:"data,omitempty"`
	Results    json.RawMessage `json:"results,omitempty"`
	Result     json.RawMessage `json:"result,omitempty"`
}
