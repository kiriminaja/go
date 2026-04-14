package types

type PricingExpressPayload struct {
	Origin      int      `json:"origin"`
	Destination int      `json:"destination"`
	Weight      int      `json:"weight"`
	ItemValue   any      `json:"item_value"`
	Insurance   int      `json:"insurance"`
	Courier     []string `json:"courier"`
}

type PricingInstantLocationPayload struct {
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
	Address string  `json:"address"`
}

type InstantService string

const (
	InstantServiceGrabExpress InstantService = "grab_express"
	InstantServiceBorzo       InstantService = "borzo"
	InstantServiceGosend      InstantService = "gosend"
)

type InstantVehicle string

const (
	InstantVehicleBike InstantVehicle = "motor"
	InstantVehicleCar  InstantVehicle = "mobil"
)

type PricingInstantPayload struct {
	Service     []InstantService              `json:"service"`
	ItemPrice   float64                       `json:"item_price"`
	Origin      PricingInstantLocationPayload `json:"origin"`
	Destination PricingInstantLocationPayload `json:"destination"`
	Weight      int                           `json:"weight"`
	Vehicle     InstantVehicle                `json:"vehicle"`
	Timezone    string                        `json:"timezone"`
}

type ExpressService = string

const (
	ExpressServiceTiki         ExpressService = "tiki"
	ExpressServicePos          ExpressService = "posindonesia"
	ExpressServicePaxel        ExpressService = "paxel"
	ExpressServiceNinja        ExpressService = "ninja"
	ExpressServiceRPX          ExpressService = "rpx"
	ExpressServiceLionParcel   ExpressService = "lion"
	ExpressServiceJTCargo      ExpressService = "jtcargo"
	ExpressServiceSentralCargo ExpressService = "sentral"
	ExpressServiceAnterAja     ExpressService = "anteraja"
	ExpressServiceNCS          ExpressService = "ncs"
	ExpressServiceSicepat      ExpressService = "sicepat"
	ExpressServiceSAP          ExpressService = "sap"
	ExpressServiceIDExpress    ExpressService = "idx"
	ExpressServiceJNE          ExpressService = "jne"
	ExpressServiceJNT          ExpressService = "jnt"
	ExpressServiceSPX          ExpressService = "spx"
)
