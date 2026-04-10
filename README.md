# KiriminAja Go SDK

Official Go SDK for the [KiriminAja](https://kiriminaja.com) shipping API.

## Installation

```bash
go get github.com/kiriminaja/go
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	kiriminaja "github.com/kiriminaja/go"
)

func main() {
	client := kiriminaja.New(kiriminaja.Config{
		Env:    kiriminaja.EnvSandbox,
		APIKey: "your-api-key",
	})

	provinces, err := client.Address.Provinces()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(provinces)
}
```

## Configuration

```go
client := kiriminaja.New(kiriminaja.Config{
	Env:     kiriminaja.EnvProduction,  // or kiriminaja.EnvSandbox (default)
	APIKey:  "your-api-key",
	BaseURL: "",                         // optional, overrides env-based URL
	HTTPClient: &http.Client{},          // optional, uses http.DefaultClient
})
```

## Services

### Address

```go
client.Address.Provinces()
client.Address.Cities(provinceID)
client.Address.Districts(cityID)
client.Address.SubDistricts(districtID)
client.Address.DistrictsByName("jakarta")
```

### Courier

```go
client.Courier.List()
client.Courier.Group()
client.Courier.Detail("jne")
client.Courier.SetWhitelistServices([]string{"jne_reg", "jne_yes"})
```

### Coverage Area

```go
client.CoverageArea.PricingExpress(kiriminaja.PricingExpressPayload{...})
client.CoverageArea.PricingInstant(kiriminaja.PricingInstantPayload{...})
```

### Order

```go
// Express
client.Order.Express.Track("order-id")
client.Order.Express.Cancel("awb", "reason")
client.Order.Express.RequestPickupV5(payload)
client.Order.Express.RequestPickupV61(payload)

// Instant
client.Order.Instant.Create(payload)
client.Order.Instant.Track("order-id")
client.Order.Instant.Cancel("order-id")
client.Order.Instant.FindNewDriver("order-id")
```

### Payment

```go
client.Payment.GetPayment("payment-id")
```

### Pickup

```go
client.Pickup.Schedules()
```
