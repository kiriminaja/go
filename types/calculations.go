package types

type CalculateCODDataItem struct {
	CourierCode        string `json:"courier_code"`
	CourierServiceCode string `json:"courier_service_code"`
	DiscountAmount     int    `json:"discount_amount,omitempty"`
	InsuranceAmount    int    `json:"insurance_amount,omitempty"`
	ShippingCost       int    `json:"shipping_cost,omitempty"`
}

type CalculateCODRequest struct {
	ItemPrice                  int                     `json:"item_price"`
	Data                       []CalculateCODDataItem  `json:"data"`
	CustomCOD                  int                     `json:"custom_cod,omitempty"`
	ExcludeCODAmountValidation bool                    `json:"exclude_cod_amount_validation,omitempty"`
}

type CalculateCODMessage struct {
	MessageType string `json:"MessageType"`
	Message     string `json:"message"`
}

type CalculateCODResult struct {
	BillableAmount      string               `json:"billable_amount"`
	CourierCode         string               `json:"courier_code"`
	CourierServiceCode  string               `json:"courier_service_code"`
	Fee                 string               `json:"fee"`
	FeePercentage       float64              `json:"fee_percentage"`
	IsSupportCOD        bool                 `json:"is_support_cod"`
	Message             CalculateCODMessage  `json:"message"`
	MinimumCustomCOD    string               `json:"minimum_custom_cod"`
	MinimumFee          string               `json:"minimum_fee"`
	TaxAmount           string               `json:"tax_amount"`
	TaxPercentage       float64              `json:"tax_percentage"`
	TotalFee            string               `json:"total_fee"`
	WithdrawalAmount    string               `json:"withdrawal_amount"`
}

type CalculateCODResponse struct {
	Status  bool                  `json:"status"`
	Text    string                `json:"text"`
	Method  string                `json:"method"`
	Code    string                `json:"code"`
	Results []CalculateCODResult  `json:"results"`
}
