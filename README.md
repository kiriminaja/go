# KiriminAja Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/kiriminaja/go.svg)](https://pkg.go.dev/github.com/kiriminaja/go)
[![Latest Version](https://img.shields.io/github/v/tag/kiriminaja/go?label=version&sort=semver)](https://github.com/kiriminaja/go/releases)
[![license](https://img.shields.io/github/license/kiriminaja/go)](LICENSE)

Official Go SDK for the [KiriminAja](https://kiriminaja.com) logistics API. Zero dependencies beyond the standard library.

## Requirements

- Go 1.21+

## Installation

```bash
go get github.com/kiriminaja/go
```

---

## Quick Start

Create a client with `kiriminaja.New()`, then call any service method.

```go
package main

import (
    "fmt"
    "log"
    "os"

    kiriminaja "github.com/kiriminaja/go"
)

func main() {
    client := kiriminaja.New(kiriminaja.Config{
        Env:    kiriminaja.EnvSandbox, // or kiriminaja.EnvProduction
        APIKey: os.Getenv("KIRIMINAJA_API_KEY"),
    })

    // Use any service
    provinces, err := client.Address.Provinces()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Provinces: %+v\n", provinces)
}
```

---

## Config Options

| Option       | Type           | Default              | Description                            |
| ------------ | -------------- | -------------------- | -------------------------------------- |
| `Env`        | `Env`          | `EnvSandbox`         | Target environment                     |
| `APIKey`     | `string`       | —                    | Your KiriminAja API key                |
| `BaseURL`    | `string`       | Derived from `Env`   | Override the base URL                  |
| `HTTPClient` | `*http.Client` | `http.DefaultClient` | Custom HTTP client (proxy / test mock) |

```go
// Custom base URL
client := kiriminaja.New(kiriminaja.Config{
    BaseURL: "https://tdev.kiriminaja.com",
    APIKey:  os.Getenv("KIRIMINAJA_API_KEY"),
})

// Custom HTTP client (e.g. with timeout or transport)
client := kiriminaja.New(kiriminaja.Config{
    APIKey: "...",
    HTTPClient: &http.Client{Timeout: 10 * time.Second},
})
```

---

## Services

### Address

```go
// List all provinces
client.Address.Provinces()

// Cities in a province (provinsi_id)
client.Address.Cities(5)

// Districts in a city (kabupaten_id)
client.Address.Districts(12)

// Sub-districts in a district (kecamatan_id)
client.Address.SubDistricts(77)

// Search districts by name
client.Address.DistrictsByName("jakarta")
```

---

### Coverage Area & Pricing

```go
// Express shipping rates
client.CoverageArea.PricingExpress(types.PricingExpressPayload{
    Origin:      1,
    Destination: 2,
    Weight:      1000, // grams
    ItemValue:   50000,
    Insurance:   0,
    Courier:     []string{"jne", "jnt"},
})

// Instant (same-day) rates
client.CoverageArea.PricingInstant(types.PricingInstantPayload{
    Service:     []types.InstantService{types.InstantServiceGosend},
    ItemPrice:   10000,
    Origin:      types.PricingInstantLocationPayload{Lat: -6.2, Long: 106.8, Address: "Jl. Sudirman No.1"},
    Destination: types.PricingInstantLocationPayload{Lat: -6.21, Long: 106.81, Address: "Jl. Thamrin No.5"},
    Weight:      1000,
    Vehicle:     types.InstantVehicleBike,
    Timezone:    "Asia/Jakarta",
})
```

---

### Order — Express

```go
// Track by order ID
client.Order.Express.Track("ORDER123")

// Cancel by AWB
client.Order.Express.Cancel("AWB123456", "Customer request")

// Request pickup
client.Order.Express.RequestPickup(types.RequestPickupPayload{
    Address:     "Jl. Jodipati No.29",
    Phone:       "08133345678",
    Name:        "Tokotries",
    KecamatanID: 548,
    Schedule:    "2021-11-30 22:00:00",
    Packages: []types.RequestPickupPackage{
        {
            OrderID:                "YGL-000000019",
            DestinationName:        "Flag Test",
            DestinationPhone:       "082223323333",
            DestinationAddress:     "Jl. Magelang KM 11",
            DestinationKecamatanID: 548,
            Weight:                 520,
            Width:                  8,
            Length:                 8,
            Height:                 8,
            ItemValue:              275000,
            ShippingCost:           65000,
            Service:                "jne",
            ServiceType:            "REG23",
            COD:                    0,
            PackageTypeID:          7,
            ItemName:               "TEST Item name",
        },
    },
})
```

---

### Order — Instant

```go
// Create instant pickup
client.Order.Instant.Create(types.InstantPickupPayload{
    Service:     types.InstantServiceGosend,
    ServiceType: "instant",
    Vehicle:     types.InstantVehicleBike,
    OrderPrefix: "BDI",
    Packages: []types.InstantPickupPackage{
        {
            OriginName:             "Rizky",
            OriginPhone:            "081280045616",
            OriginLat:              -7.854584,
            OriginLong:             110.331154,
            OriginAddress:          "Wirobrajan, Yogyakarta",
            OriginAddressNote:      "Dekat Kantor",
            DestinationName:        "Okka",
            DestinationPhone:       "081280045616",
            DestinationLat:         -7.776192,
            DestinationLong:        110.325053,
            DestinationAddress:     "Godean, Sleman",
            DestinationAddressNote: "Dekat Pasar",
            ShippingPrice:          34000,
            Item: types.InstantPickupItem{
                Name:        "Barang 1",
                Description: "Barang 1 Description",
                Price:       20000,
                Weight:      1000,
            },
        },
    },
})

// Find a new driver for an existing order
client.Order.Instant.FindNewDriver("ORDER123")

// Cancel instant order
client.Order.Instant.Cancel("ORDER123")

// Track instant order
client.Order.Instant.Track("ORDER123")
```

---

### Courier

```go
// List available couriers
client.Courier.List()

// Courier groups
client.Courier.Group()

// Courier service detail
client.Courier.Detail("jne")

// Set whitelist services
client.Courier.SetWhitelistServices([]string{"jne_reg", "jne_yes"})
```

---

### Pickup Schedules

```go
client.Pickup.Schedules()
```

---

### Payment

```go
client.Payment.GetPayment("PAY123")
```

---

## Development

```bash
go build ./...   # build all packages
go test ./...    # run tests
```
