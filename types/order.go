package types

type RequestPickupPackage struct {
	OrderID                string  `json:"order_id"`
	DestinationName        string  `json:"destination_name"`
	DestinationPhone       string  `json:"destination_phone"`
	DestinationAddress     string  `json:"destination_address"`
	DestinationKecamatanID int     `json:"destination_kecamatan_id"`
	DestinationKelurahanID int     `json:"destination_kelurahan_id,omitempty"`
	DestinationZipcode     string  `json:"destination_zipcode,omitempty"`
	Weight                 int     `json:"weight"`
	Width                  int     `json:"width"`
	Length                 int     `json:"length"`
	Height                 int     `json:"height"`
	Qty                    int     `json:"qty,omitempty"`
	ItemValue              int     `json:"item_value"`
	ShippingCost           int     `json:"shipping_cost"`
	Service                string  `json:"service"`
	ServiceType            string  `json:"service_type"`
	InsuranceAmount        float64 `json:"insurance_amount,omitempty"`
	COD                    int     `json:"cod"`
	PackageTypeID          int     `json:"package_type_id"`
	ItemName               string  `json:"item_name"`
	Drop                   bool    `json:"drop,omitempty"`
	Note                   string  `json:"note,omitempty"`
}

type RequestPickupPayload struct {
	Address      string                 `json:"address"`
	Phone        string                 `json:"phone"`
	Name         string                 `json:"name"`
	Zipcode      string                 `json:"zipcode,omitempty"`
	KecamatanID  int                    `json:"kecamatan_id"`
	KelurahanID  int                    `json:"kelurahan_id,omitempty"`
	Latitude     float64                `json:"latitude,omitempty"`
	Longitude    float64                `json:"longitude,omitempty"`
	Packages     []RequestPickupPackage `json:"packages"`
	Schedule     string                 `json:"schedule"`
	PlatformName string                 `json:"platform_name,omitempty"`
}

type InstantPickupItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Weight      int    `json:"weight"`
}

type InstantPickupPackage struct {
	OriginName             string            `json:"origin_name"`
	OriginPhone            string            `json:"origin_phone"`
	OriginLat              float64           `json:"origin_lat"`
	OriginLong             float64           `json:"origin_long"`
	OriginAddress          string            `json:"origin_address"`
	OriginAddressNote      string            `json:"origin_address_note"`
	DestinationName        string            `json:"destination_name"`
	DestinationPhone       string            `json:"destination_phone"`
	DestinationLat         float64           `json:"destination_lat"`
	DestinationLong        float64           `json:"destination_long"`
	DestinationAddress     string            `json:"destination_address"`
	DestinationAddressNote string            `json:"destination_address_note"`
	ShippingPrice          int               `json:"shipping_price"`
	Item                   InstantPickupItem `json:"item"`
}

type InstantPickupPayload struct {
	Service     InstantService         `json:"service"`
	ServiceType string                 `json:"service_type"`
	Vehicle     InstantVehicle         `json:"vehicle"`
	OrderPrefix string                 `json:"order_prefix"`
	Packages    []InstantPickupPackage `json:"packages"`
}
