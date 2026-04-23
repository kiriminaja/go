package tests

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	kiriminaja "github.com/kiriminaja/go"
	kahttp "github.com/kiriminaja/go/http"
	"github.com/kiriminaja/go/config"
	"github.com/kiriminaja/go/types"
)

type requestCall struct {
	Method string
	URL    string
	Body   string
	Header http.Header
}

type mockTransport struct {
	calls    []requestCall
	response string
	status   int
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var body string
	if req.Body != nil {
		data, err := io.ReadAll(req.Body)
		if err == nil {
			body = string(data)
		}
	}
	m.calls = append(m.calls, requestCall{
		Method: req.Method,
		URL:    req.URL.String(),
		Body:   body,
		Header: req.Header.Clone(),
	})
	status := m.status
	if status == 0 {
		status = 200
	}
	resp := m.response
	if resp == "" {
		resp = `{"status":true}`
	}
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(resp)),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
	}, nil
}

func newMockClient(env config.Env, apiKey string) (*kiriminaja.Client, *mockTransport) {
	transport := &mockTransport{}
	client := kiriminaja.New(kiriminaja.Config{
		Env:    env,
		APIKey: apiKey,
		HTTPClient: &http.Client{
			Transport: transport,
		},
	})
	return client, transport
}

func newMockClientWithResponse(env config.Env, status int, response string) (*kiriminaja.Client, *mockTransport) {
	transport := &mockTransport{status: status, response: response}
	client := kiriminaja.New(kiriminaja.Config{
		Env:    env,
		HTTPClient: &http.Client{
			Transport: transport,
		},
	})
	return client, transport
}

func assertContains(t *testing.T, s, substr string) {
	t.Helper()
	if !strings.Contains(s, substr) {
		t.Errorf("expected %q to contain %q", s, substr)
	}
}

func assertStartsWith(t *testing.T, s, prefix string) {
	t.Helper()
	if !strings.HasPrefix(s, prefix) {
		t.Errorf("expected %q to start with %q", s, prefix)
	}
}

func assertEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSandboxBaseURL(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.Provinces(ctx)
	if len(transport.calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(transport.calls))
	}
	assertStartsWith(t, transport.calls[0].URL, config.EnvURL[kiriminaja.EnvSandbox])
}

func TestProductionBaseURL(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvProduction, "")
	client.Address.Provinces(ctx)
	if len(transport.calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(transport.calls))
	}
	assertStartsWith(t, transport.calls[0].URL, config.EnvURL[kiriminaja.EnvProduction])
}

func TestBearerToken(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "abc")
	client.Address.Provinces(ctx)
	if len(transport.calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(transport.calls))
	}
	assertEqual(t, transport.calls[0].Header.Get("Authorization"), "Bearer abc")
	assertEqual(t, transport.calls[0].Header.Get("Accept"), "application/json")
}

func TestProvincesEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.Provinces(ctx)
	assertContains(t, transport.calls[0].URL, "/api/mitra/province")
	assertEqual(t, transport.calls[0].Method, "POST")
}

func TestCitiesEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.Cities(ctx, 5)
	assertContains(t, transport.calls[0].URL, "/api/mitra/city")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(map[string]any{"provinsi_id": 5})
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestDistrictsEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.Districts(ctx, 12)
	assertContains(t, transport.calls[0].URL, "/api/mitra/kecamatan")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(map[string]any{"kabupaten_id": 12})
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestSubDistrictsEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.SubDistricts(ctx, 77)
	assertContains(t, transport.calls[0].URL, "/api/mitra/kelurahan")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(map[string]any{"kecamatan_id": 77})
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestDistrictsByNameEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.DistrictsByName(ctx, "jakarta")
	assertContains(t, transport.calls[0].URL, "/api/mitra/v2/get_address_by_name")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	assertEqual(t, transport.calls[0].Body, `{"search":"jakarta"}`)
}

func TestPricingExpressEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	payload := types.PricingExpressPayload{
		Origin:      1,
		Destination: 2,
		Weight:      1000,
		ItemValue:   5000,
		Insurance:   0,
		Courier:     []types.ExpressService{types.ExpressServiceJNE, "other"},
	}
	client.CoverageArea.PricingExpress(ctx, payload)
	assertContains(t, transport.calls[0].URL, "/api/mitra/v6.1/shipping_price")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(payload)
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestPricingInstantEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	payload := types.PricingInstantPayload{
		Service:     []types.InstantService{types.InstantServiceGrabExpress, "other"},
		ItemPrice:   10000,
		Origin:      types.PricingInstantLocationPayload{Lat: -6.2, Long: 106.8, Address: "A"},
		Destination: types.PricingInstantLocationPayload{Lat: -6.21, Long: 106.81, Address: "B"},
		Weight:      1000,
		Vehicle:     types.InstantVehicleBike,
		Timezone:    "Asia/Jakarta",
	}
	client.CoverageArea.PricingInstant(ctx, payload)
	assertContains(t, transport.calls[0].URL, "/api/mitra/v4/instant/pricing")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(payload)
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestExpressCancelEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Order.Express.Cancel(ctx, "AWB123", "reason here")
	u, _ := url.Parse(transport.calls[0].URL)
	assertContains(t, u.Path, "/api/mitra/v3/cancel_shipment")
	assertEqual(t, u.Query().Get("awb"), "AWB123")
	assertEqual(t, u.Query().Get("reason"), "reason here")
	assertEqual(t, transport.calls[0].Method, "POST")
}

func TestExpressTrackEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Order.Express.Track(ctx, "OID_EXP_1")
	assertContains(t, transport.calls[0].URL, "/api/mitra/tracking")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	assertEqual(t, transport.calls[0].Body, `{"order_id":"OID_EXP_1"}`)
}

func TestExpressRequestPickupEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	payload := types.RequestPickupPayload{
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
				Items: []types.RequestPickupItem{
					{
						Name:   "Kaos Polos",
						Price:  125000,
						Qty:    2,
						Weight: 260,
						Width:  4,
						Length: 4,
						Height: 4,
						Metadata: &types.RequestPickupItemMetadata{
							SKU:          "KP-001",
							VariantLabel: "Merah / L",
						},
					},
				},
			},
		},
	}
	client.Order.Express.RequestPickup(ctx, payload)
	assertContains(t, transport.calls[0].URL, "/api/mitra/v6.1/request_pickup")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(payload)
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestInstantTrackEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Order.Instant.Track(ctx, "OID123")
	assertContains(t, transport.calls[0].URL, "/api/mitra/v4/instant/tracking/OID123")
	assertEqual(t, transport.calls[0].Method, "GET")
}

func TestInstantCreateEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	payload := types.InstantPickupPayload{
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
	}
	client.Order.Instant.Create(ctx, payload)
	assertContains(t, transport.calls[0].URL, "/api/mitra/v4/instant/pickup/request")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(payload)
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestInstantFindNewDriverEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Order.Instant.FindNewDriver(ctx, "OID123")
	assertContains(t, transport.calls[0].URL, "/api/mitra/v4/instant/pickup/find-new-driver")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	assertEqual(t, transport.calls[0].Body, `{"order_id":"OID123"}`)
}

func TestInstantCancelEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Order.Instant.Cancel(ctx, "OID123")
	assertContains(t, transport.calls[0].URL, "/api/mitra/v4/instant/pickup/void/OID123")
	assertEqual(t, transport.calls[0].Method, "DELETE")
}

func TestCourierListEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Courier.List(ctx)
	assertContains(t, transport.calls[0].URL, "/api/mitra/couriers")
	assertEqual(t, transport.calls[0].Method, "POST")
}

func TestCourierGroupEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Courier.Group(ctx)
	assertContains(t, transport.calls[0].URL, "/api/mitra/couriers_group")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Body, "")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "")
}

func TestCourierDetailEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Courier.Detail(ctx, "jne")
	assertContains(t, transport.calls[0].URL, "/api/mitra/courier_services")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	assertEqual(t, transport.calls[0].Body, `{"courier_code":"jne"}`)
}

func TestCourierSetWhitelistServicesEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Courier.SetWhitelistServices(ctx, []string{"jne_reg", "jne_yes"})
	assertContains(t, transport.calls[0].URL, "/api/mitra/v3/set_whitelist_services")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(map[string]any{"services": []string{"jne_reg", "jne_yes"}})
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestPickupSchedulesEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Pickup.Schedules(ctx)
	assertContains(t, transport.calls[0].URL, "/api/mitra/v2/schedules")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Body, "")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "")
}

func TestPaymentGetPaymentEndpoint(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Payment.GetPayment(ctx, "PAY123")
	assertContains(t, transport.calls[0].URL, "/api/mitra/v2/get_payment")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	assertEqual(t, transport.calls[0].Body, `{"payment_id":"PAY123"}`)
}

func TestCoverageAreaDelegatesAddressMethods(t *testing.T) {
	ctx := context.Background()
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.CoverageArea.Provinces(ctx)
	assertContains(t, transport.calls[0].URL, "/api/mitra/province")
	client.CoverageArea.Cities(ctx, 1)
	assertContains(t, transport.calls[1].URL, "/api/mitra/city")
	client.CoverageArea.Districts(ctx, 1)
	assertContains(t, transport.calls[2].URL, "/api/mitra/kecamatan")
	client.CoverageArea.SubDistricts(ctx, 1)
	assertContains(t, transport.calls[3].URL, "/api/mitra/kelurahan")
	client.CoverageArea.DistrictsByName(ctx, "test")
	assertContains(t, transport.calls[4].URL, "/api/mitra/v2/get_address_by_name")
}

// --- Negative tests ---

func TestValidationEmptyOrderID(t *testing.T) {
	ctx := context.Background()
	client, _ := newMockClient(kiriminaja.EnvSandbox, "")
	_, err := client.Order.Express.Track(ctx, "")
	if err == nil {
		t.Error("expected error for empty orderID, got nil")
	}
}

func TestValidationEmptyAWB(t *testing.T) {
	ctx := context.Background()
	client, _ := newMockClient(kiriminaja.EnvSandbox, "")
	_, err := client.Order.Express.Cancel(ctx, "", "reason")
	if err == nil {
		t.Error("expected error for empty awb, got nil")
	}
}

func TestValidationEmptyInstantOrderID(t *testing.T) {
	ctx := context.Background()
	client, _ := newMockClient(kiriminaja.EnvSandbox, "")
	_, err := client.Order.Instant.Track(ctx, "")
	if err == nil {
		t.Error("expected error for empty orderID, got nil")
	}
	_, err = client.Order.Instant.Cancel(ctx, "")
	if err == nil {
		t.Error("expected error for empty orderID, got nil")
	}
	_, err = client.Order.Instant.FindNewDriver(ctx, "")
	if err == nil {
		t.Error("expected error for empty orderID, got nil")
	}
}

func TestValidationEmptyPaymentID(t *testing.T) {
	ctx := context.Background()
	client, _ := newMockClient(kiriminaja.EnvSandbox, "")
	_, err := client.Payment.GetPayment(ctx, "")
	if err == nil {
		t.Error("expected error for empty paymentID, got nil")
	}
}

func TestValidationEmptyCourierCode(t *testing.T) {
	ctx := context.Background()
	client, _ := newMockClient(kiriminaja.EnvSandbox, "")
	_, err := client.Courier.Detail(ctx, "")
	if err == nil {
		t.Error("expected error for empty courierCode, got nil")
	}
}

func TestValidationEmptyServices(t *testing.T) {
	ctx := context.Background()
	client, _ := newMockClient(kiriminaja.EnvSandbox, "")
	_, err := client.Courier.SetWhitelistServices(ctx, nil)
	if err == nil {
		t.Error("expected error for nil services, got nil")
	}
}

func TestValidationZeroProvinsiID(t *testing.T) {
	ctx := context.Background()
	client, _ := newMockClient(kiriminaja.EnvSandbox, "")
	_, err := client.Address.Cities(ctx, 0)
	if err == nil {
		t.Error("expected error for zero provinsiID, got nil")
	}
}

func TestValidationEmptySearchQuery(t *testing.T) {
	ctx := context.Background()
	client, _ := newMockClient(kiriminaja.EnvSandbox, "")
	_, err := client.Address.DistrictsByName(ctx, "")
	if err == nil {
		t.Error("expected error for empty search, got nil")
	}
}

func TestAPIErrorOnNon2xx(t *testing.T) {
	ctx := context.Background()
	client, _ := newMockClientWithResponse(kiriminaja.EnvSandbox, 401, `{"message":"Unauthorized"}`)
	_, err := client.Address.Provinces(ctx)
	if err == nil {
		t.Fatal("expected error for 401 response, got nil")
	}
	apiErr, ok := err.(*kahttp.APIError)
	if !ok {
		t.Fatalf("expected *kahttp.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 401 {
		t.Errorf("expected StatusCode 401, got %d", apiErr.StatusCode)
	}
	assertContains(t, apiErr.Body, "Unauthorized")
}

func TestAPIErrorOn500(t *testing.T) {
	ctx := context.Background()
	client, _ := newMockClientWithResponse(kiriminaja.EnvSandbox, 500, `<html>Internal Server Error</html>`)
	_, err := client.Address.Provinces(ctx)
	if err == nil {
		t.Fatal("expected error for 500 response, got nil")
	}
	apiErr, ok := err.(*kahttp.APIError)
	if !ok {
		t.Fatalf("expected *kahttp.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 500 {
		t.Errorf("expected StatusCode 500, got %d", apiErr.StatusCode)
	}
}

func TestMalformedJSONResponse(t *testing.T) {
	ctx := context.Background()
	client, _ := newMockClientWithResponse(kiriminaja.EnvSandbox, 200, `not-json`)
	_, err := client.Address.Provinces(ctx)
	if err == nil {
		t.Error("expected error for malformed JSON, got nil")
	}
}
