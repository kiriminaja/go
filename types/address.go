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

type PricingInstantPayload struct {
	Service     []string                      `json:"service"`
	ItemPrice   float64                       `json:"item_price"`
	Origin      PricingInstantLocationPayload `json:"origin"`
	Destination PricingInstantLocationPayload `json:"destination"`
	Weight      int                           `json:"weight"`
	Vehicle     string                        `json:"vehicle"`
	Timezone    string                        `json:"timezone"`
}
