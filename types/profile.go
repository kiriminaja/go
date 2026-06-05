package types

type ProfileMeta struct {
	HasPin        bool   `json:"has_pin"`
	PaymentMethod string `json:"payment_method"`
}

type ProfileData struct {
	ID       int         `json:"id"`
	Email    string      `json:"email"`
	Name     string      `json:"name"`
	Status   string      `json:"status"`
	Metadata ProfileMeta `json:"metadata"`
}

type ProfileResponse struct {
	Status  bool        `json:"status"`
	Text    string      `json:"text"`
	Method  string      `json:"method"`
	Code    string      `json:"code"`
	Results ProfileData `json:"results"`
}
