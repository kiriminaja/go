package types

type CancelExpressOrderResponse struct {
	BaseResponse
	Data struct {
		Success string `json:"success"`
		Date    string `json:"date"`
	} `json:"data"`
}

type ExpressTrackingImages struct {
	CameraImg    *string `json:"camera_img"`
	SignatureImg *string `json:"signature_img"`
	PopImg       *string `json:"pop_img"`
}

type ExpressTrackingCosts struct {
	AddCost          float64 `json:"add_cost"`
	Currency         string  `json:"currency"`
	COD              float64 `json:"cod"`
	InsuranceAmount  float64 `json:"insurance_amount"`
	InsurancePercent float64 `json:"insurance_percent"`
	DiscountAmount   float64 `json:"discount_amount"`
	SubsidiAmount    float64 `json:"subsidi_amount"`
	ShippingCost     float64 `json:"shipping_cost"`
	Correction       float64 `json:"correction"`
}

type ExpressTrackingLocation struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	City    string `json:"city"`
	ZipCode any    `json:"zip_code"`
}

type ExpressTrackingDetails struct {
	AWB           *string                 `json:"awb"`
	SignatureCode *string                 `json:"signature_code"`
	SortingCode   *string                 `json:"sorting_code"`
	OrderID       string                  `json:"order_id"`
	StatusCode    *int                    `json:"status_code"`
	Estimation    string                  `json:"estimation"`
	Service       string                  `json:"service"`
	ServiceName   string                  `json:"service_name"`
	Drop          bool                    `json:"drop"`
	ShippedAt     *string                 `json:"shipped_at"`
	Delivered     bool                    `json:"delivered"`
	DeliveredAt   *string                 `json:"delivered_at"`
	Refunded      bool                    `json:"refunded"`
	RefundedAt    *string                 `json:"refunded_at"`
	Images        ExpressTrackingImages   `json:"images"`
	Costs         ExpressTrackingCosts    `json:"costs"`
	Origin        ExpressTrackingLocation `json:"origin"`
	Destination   ExpressTrackingLocation `json:"destination"`
}

type ExpressTrackingHistory struct {
	CreatedAt  string `json:"created_at"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Driver     string `json:"driver"`
	Receiver   string `json:"receiver"`
}

type ExpressTrackingResponse struct {
	BaseResponse
	Details   ExpressTrackingDetails   `json:"details"`
	Histories []ExpressTrackingHistory `json:"histories"`
}

type InstantTrackingDriver struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
	Photo *string `json:"photo"`
}

type InstantTrackingLocation struct {
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	AddressNote string  `json:"address_note"`
	Phone       string  `json:"phone"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
}

type InstantTrackingDate struct {
	CreatedAt   string  `json:"created_at"`
	FinishedAt  *string `json:"finished_at"`
	AllocatedAt *string `json:"allocated_at"`
	CanceledAt  *string `json:"canceled_at"`
}

type InstantTrackingCost struct {
	ShippingCost float64  `json:"shipping_cost"`
	Insurance    *float64 `json:"insurance"`
	AdminFee     float64  `json:"admin_fee"`
	TotalPrice   float64  `json:"total_price"`
}

type InstantTrackingItem struct {
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type InstantTrackingResult struct {
	Driver            InstantTrackingDriver   `json:"driver"`
	Origin            InstantTrackingLocation `json:"origin"`
	Destination       InstantTrackingLocation `json:"destination"`
	Date              InstantTrackingDate     `json:"date"`
	Cost              InstantTrackingCost     `json:"cost"`
	Item              InstantTrackingItem     `json:"item"`
	OrderID           string                  `json:"order_id"`
	Service           string                  `json:"service"`
	ServiceType       string                  `json:"service_type"`
	TrackingCode      string                  `json:"tracking_code"`
	CancelDescription *string                 `json:"cancel_description"`
	LiveTrackingURL   *string                 `json:"live_tracking_url"`
}

type InstantTrackingResponse struct {
	BaseResponse
	Result *InstantTrackingResult `json:"result,omitempty"`
}

type CancelInstantPayment struct {
	PaymentID  string  `json:"payment_id"`
	Amount     float64 `json:"amount"`
	StatusCode int     `json:"status_code"`
	QRContent  *string `json:"qr_content"`
	PayTime    string  `json:"pay_time"`
}

type CancelInstantPackageLocation struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Phone     string  `json:"phone"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CancelInstantPackage struct {
	AWB          string                       `json:"awb"`
	OrderID      string                       `json:"order_id"`
	Service      string                       `json:"service"`
	ServiceType  string                       `json:"service_type"`
	Status       int                          `json:"status"`
	LiveTrackURL *string                      `json:"live_track_url"`
	Polyline     string                       `json:"polyline,omitempty"`
	Origin       CancelInstantPackageLocation `json:"origin"`
	Destination  CancelInstantPackageLocation `json:"destination"`
}

type CancelInstantOrderResult struct {
	Payment  CancelInstantPayment   `json:"payment"`
	Packages []CancelInstantPackage `json:"packages"`
}

type CancelInstantOrderResponse struct {
	BaseResponse
	Result CancelInstantOrderResult `json:"result"`
}

type CreateInstantPickupResponse = KAResponse
type FindNewInstantDriverResponse = KAResponse
