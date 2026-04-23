package types

type CreditBalanceData struct {
	Balance float64 `json:"balance"`
}

type CreditBalanceResponse struct {
	Status bool              `json:"status"`
	Text   string            `json:"text"`
	Method string            `json:"method"`
	Code   string            `json:"code"`
	Data   CreditBalanceData `json:"data"`
}
