package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	kiriminaja "github.com/kiriminaja/go"
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
	calls []requestCall
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var body string
	if req.Body != nil {
		data, _ := io.ReadAll(req.Body)
		body = string(data)
	}
	m.calls = append(m.calls, requestCall{
		Method: req.Method,
		URL:    req.URL.String(),
		Body:   body,
		Header: req.Header.Clone(),
	})
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(`{"status":true}`)),
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
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.Provinces()
	if len(transport.calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(transport.calls))
	}
	assertStartsWith(t, transport.calls[0].URL, config.EnvURL[kiriminaja.EnvSandbox])
}

func TestProductionBaseURL(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvProduction, "")
	client.Address.Provinces()
	if len(transport.calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(transport.calls))
	}
	assertStartsWith(t, transport.calls[0].URL, config.EnvURL[kiriminaja.EnvProduction])
}

func TestBearerToken(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "abc")
	client.Address.Provinces()
	if len(transport.calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(transport.calls))
	}
	assertEqual(t, transport.calls[0].Header.Get("Authorization"), "Bearer abc")
	assertEqual(t, transport.calls[0].Header.Get("Accept"), "application/json")
}

func TestProvincesEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.Provinces()
	assertContains(t, transport.calls[0].URL, "/api/mitra/province")
	assertEqual(t, transport.calls[0].Method, "POST")
}

func TestCitiesEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.Cities(5)
	assertContains(t, transport.calls[0].URL, "/api/mitra/city")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(map[string]any{"provinsi_id": 5})
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestDistrictsEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.Districts(12)
	assertContains(t, transport.calls[0].URL, "/api/mitra/kecamatan")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(map[string]any{"kabupaten_id": 12})
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestSubDistrictsEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.SubDistricts(77)
	assertContains(t, transport.calls[0].URL, "/api/mitra/kelurahan")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(map[string]any{"kecamatan_id": 77})
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestDistrictsByNameEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Address.DistrictsByName("jakarta")
	assertContains(t, transport.calls[0].URL, "/api/mitra/v2/get_address_by_name")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	assertEqual(t, transport.calls[0].Body, `{"search":"jakarta"}`)
}

func TestPricingExpressEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	payload := types.PricingExpressPayload{
		Origin:      1,
		Destination: 2,
		Weight:      1000,
		ItemValue:   5000,
		Insurance:   0,
		Courier:     []string{"jne"},
	}
	client.CoverageArea.PricingExpress(payload)
	assertContains(t, transport.calls[0].URL, "/api/mitra/v6.1/shipping_price")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(payload)
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestPricingInstantEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	payload := types.PricingInstantPayload{
		Service:     []types.InstantService{types.InstantServiceGrabExpress},
		ItemPrice:   10000,
		Origin:      types.PricingInstantLocationPayload{Lat: -6.2, Long: 106.8, Address: "A"},
		Destination: types.PricingInstantLocationPayload{Lat: -6.21, Long: 106.81, Address: "B"},
		Weight:      1000,
		Vehicle:     types.InstantVehicleMotor,
		Timezone:    "Asia/Jakarta",
	}
	client.CoverageArea.PricingInstant(payload)
	assertContains(t, transport.calls[0].URL, "/api/mitra/v4/instant/pricing")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(payload)
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestExpressCancelEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Order.Express.Cancel("AWB123", "reason here")
	u, _ := url.Parse(transport.calls[0].URL)
	assertContains(t, u.Path, "/api/mitra/v3/cancel_shipment")
	assertEqual(t, u.Query().Get("awb"), "AWB123")
	assertEqual(t, u.Query().Get("reason"), "reason here")
	assertEqual(t, transport.calls[0].Method, "POST")
}

func TestExpressTrackEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Order.Express.Track("OID_EXP_1")
	assertContains(t, transport.calls[0].URL, "/api/mitra/tracking")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	assertEqual(t, transport.calls[0].Body, `{"order_id":"OID_EXP_1"}`)
}

func TestExpressRequestPickupEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	payload := map[string]any{"foo": "bar"}
	client.Order.Express.RequestPickup(payload)
	assertContains(t, transport.calls[0].URL, "/api/mitra/v6.1/request_pickup")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(payload)
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestInstantTrackEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Order.Instant.Track("OID123")
	assertContains(t, transport.calls[0].URL, "/api/mitra/v4/instant/tracking/OID123")
	assertEqual(t, transport.calls[0].Method, "GET")
}

func TestInstantCreateEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	payload := map[string]any{
		"origin":      map[string]any{"address": "A"},
		"destination": map[string]any{"address": "B"},
	}
	client.Order.Instant.Create(payload)
	assertContains(t, transport.calls[0].URL, "/api/mitra/v4/instant/pickup/request")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(payload)
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestInstantFindNewDriverEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Order.Instant.FindNewDriver("OID123")
	assertContains(t, transport.calls[0].URL, "/api/mitra/v4/instant/pickup/find-new-driver")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	assertEqual(t, transport.calls[0].Body, `{"order_id":"OID123"}`)
}

func TestInstantCancelEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Order.Instant.Cancel("OID123")
	assertContains(t, transport.calls[0].URL, "/api/mitra/v4/instant/pickup/void/OID123")
	assertEqual(t, transport.calls[0].Method, "DELETE")
}

func TestCourierListEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Courier.List()
	assertContains(t, transport.calls[0].URL, "/api/mitra/couriers")
	assertEqual(t, transport.calls[0].Method, "POST")
}

func TestCourierGroupEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Courier.Group()
	assertContains(t, transport.calls[0].URL, "/api/mitra/couriers_group")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Body, "")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "")
}

func TestCourierDetailEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Courier.Detail("jne")
	assertContains(t, transport.calls[0].URL, "/api/mitra/courier_services")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	assertEqual(t, transport.calls[0].Body, `{"courier_code":"jne"}`)
}

func TestCourierSetWhitelistServicesEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Courier.SetWhitelistServices([]string{"jne_reg", "jne_yes"})
	assertContains(t, transport.calls[0].URL, "/api/mitra/v3/set_whitelist_services")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	expected, _ := json.Marshal(map[string]any{"services": []string{"jne_reg", "jne_yes"}})
	assertEqual(t, transport.calls[0].Body, string(expected))
}

func TestPickupSchedulesEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Pickup.Schedules()
	assertContains(t, transport.calls[0].URL, "/api/mitra/v2/schedules")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Body, "")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "")
}

func TestPaymentGetPaymentEndpoint(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.Payment.GetPayment("PAY123")
	assertContains(t, transport.calls[0].URL, "/api/mitra/v2/get_payment")
	assertEqual(t, transport.calls[0].Method, "POST")
	assertEqual(t, transport.calls[0].Header.Get("Content-Type"), "application/json")
	assertEqual(t, transport.calls[0].Body, `{"payment_id":"PAY123"}`)
}

func TestCoverageAreaInheritsAddressMethods(t *testing.T) {
	client, transport := newMockClient(kiriminaja.EnvSandbox, "")
	client.CoverageArea.Provinces()
	assertContains(t, transport.calls[0].URL, "/api/mitra/province")
	client.CoverageArea.Cities(1)
	assertContains(t, transport.calls[1].URL, "/api/mitra/city")
	client.CoverageArea.Districts(1)
	assertContains(t, transport.calls[2].URL, "/api/mitra/kecamatan")
	client.CoverageArea.SubDistricts(1)
	assertContains(t, transport.calls[3].URL, "/api/mitra/kelurahan")
	client.CoverageArea.DistrictsByName("test")
	assertContains(t, transport.calls[4].URL, "/api/mitra/v2/get_address_by_name")
}
