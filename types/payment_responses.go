package types

type PaymentExpressData struct {
	PaymentID  string  `json:"payment_id"`
	QRContent  string  `json:"qr_content"`
	Method     string  `json:"method"`
	PayTime    string  `json:"pay_time"`
	Status     string  `json:"status"`
	StatusCode string  `json:"status_code"`
	Amount     float64 `json:"amount"`
	PaidAt     *string `json:"paid_at"`
	CreatedAt  string  `json:"created_at"`
}

type PaymentInstantPackageLocation struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Phone     string  `json:"phone"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PaymentInstantPackage struct {
	AWB          string                        `json:"awb"`
	OrderID      string                        `json:"order_id"`
	Service      string                        `json:"service"`
	ServiceName  string                        `json:"service_name"`
	Status       int                           `json:"status"`
	LiveTrackURL *string                       `json:"live_track_url"`
	Origin       PaymentInstantPackageLocation `json:"origin"`
	Destination  PaymentInstantPackageLocation `json:"destination"`
}

type PaymentInstantResult struct {
	PaymentID  string                  `json:"payment_id"`
	Amount     float64                 `json:"amount"`
	StatusCode int                     `json:"status_code"`
	QRContent  *string                 `json:"qr_content"`
	PayTime    *string                 `json:"pay_time"`
	Packages   []PaymentInstantPackage `json:"packages"`
}

// GetPaymentResponse covers both express and instant payment types.
// Check Data (express) or Result (instant) to determine the type.
type GetPaymentResponse struct {
	BaseResponse
	Data   *PaymentExpressData   `json:"data,omitempty"`
	Result *PaymentInstantResult `json:"result,omitempty"`
}
