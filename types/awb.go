package types

type PrintAWBRequest struct {
	AWB []string `json:"awb"`
}

type PrintAWBData struct {
	URL string `json:"url"`
}

type PrintAWBResult struct {
	Data PrintAWBData `json:"data"`
}

type PrintAWBResponse struct {
	Status bool              `json:"status"`
	Text   string            `json:"text"`
	Method string            `json:"method"`
	Code   string            `json:"code"`
	Data   PrintAWBResult    `json:"data"`
	Errors []interface{}     `json:"errors,omitempty"`
}
