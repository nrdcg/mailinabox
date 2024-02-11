package mailinabox

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDNSService_GetSecondaryNameserver(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/secondary-nameserver", testHandler("dns/nameserver.json", http.MethodGet, http.StatusOK))

	names, err := client.DNS.GetSecondaryNameserver(context.Background())
	require.NoError(t, err)

	expected := []string{"ns1.example.com"}

	assert.Equal(t, expected, names)
}

func TestDNSService_AddSecondaryNameserver(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/secondary-nameserver", testHandler("dns/addnameserver.html", http.MethodPost, http.StatusOK))

	resp, err := client.DNS.AddSecondaryNameserver(context.Background(), []string{"ns2.hostingcompany.com", "ns3.hostingcompany.com"})
	require.NoError(t, err)

	assert.Equal(t, "updated DNS: example.com", resp)
}

func TestDNSService_GetZones(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/zones", testHandler("dns/zones.json", http.MethodGet, http.StatusOK))

	names, err := client.DNS.GetZones(context.Background())
	require.NoError(t, err)

	expected := []string{"example.com"}

	assert.Equal(t, expected, names)
}

func TestDNSService_GetZoneFile(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/zonefile/example.com", testHandler("dns/zonefile.json", http.MethodGet, http.StatusOK))

	names, err := client.DNS.GetZoneFile(context.Background(), "example.com")
	require.NoError(t, err)

	assert.Equal(t, "string", names)
}

func TestDNSService_GetAllRecords(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/custom", testHandler("dns/records.json", http.MethodGet, http.StatusOK))

	records, err := client.DNS.GetAllRecords(context.Background())
	require.NoError(t, err)

	expected := []Record{
		{Name: "example.com", Type: "MX", Value: "10 example.com."},
	}

	assert.Equal(t, expected, records)
}

func TestDNSService_GetRecords(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/custom/example.com/MX", testHandler("dns/records.json", http.MethodGet, http.StatusOK))

	records, err := client.DNS.GetRecords(context.Background(), "example.com", "MX")
	require.NoError(t, err)

	expected := []Record{
		{Name: "example.com", Type: "MX", Value: "10 example.com."},
	}

	assert.Equal(t, expected, records)
}

func TestDNSService_AddRecord(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/custom/example.com/MX", testHandler("dns/addrecord.html", http.MethodPost, http.StatusOK))

	resp, err := client.DNS.AddRecord(context.Background(), Record{Name: "example.com", Type: "MX", Value: "10 example.com."})
	require.NoError(t, err)

	assert.Equal(t, "updated DNS: example.com", resp)
}

func TestDNSService_UpdateRecord(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/custom/example.com/MX", testHandler("dns/updaterecord.html", http.MethodPut, http.StatusOK))

	resp, err := client.DNS.UpdateRecord(context.Background(), Record{Name: "example.com", Type: "MX"}, "1.2.3.4")
	require.NoError(t, err)

	assert.Equal(t, "updated DNS: example.com", resp)
}

func TestDNSService_RemoveRecord(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/custom/example.com/MX", testHandler("dns/removerecord.html", http.MethodDelete, http.StatusOK))

	resp, err := client.DNS.RemoveRecord(context.Background(), Record{Name: "example.com", Type: "MX", Value: "1.2.3.4"})
	require.NoError(t, err)

	assert.Equal(t, "updated DNS: example.com", resp)
}

func TestDNSService_GetARecords(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/custom/example.com", testHandler("dns/records.json", http.MethodGet, http.StatusOK))

	records, err := client.DNS.GetARecords(context.Background(), "example.com")
	require.NoError(t, err)

	expected := []Record{
		{Name: "example.com", Type: "MX", Value: "10 example.com."},
	}

	assert.Equal(t, expected, records)
}

func TestDNSService_AddARecord(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/custom/example.com", testHandler("dns/addrecord.html", http.MethodPost, http.StatusOK))

	resp, err := client.DNS.AddARecord(context.Background(), "example.com", "1.2.3.4")
	require.NoError(t, err)

	assert.Equal(t, "updated DNS: example.com", resp)
}

func TestDNSService_UpdateARecord(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/custom/example.com", testHandler("dns/updaterecord.html", http.MethodPut, http.StatusOK))

	resp, err := client.DNS.UpdateARecord(context.Background(), "example.com", "1.2.3.4")
	require.NoError(t, err)

	assert.Equal(t, "updated DNS: example.com", resp)
}

func TestDNSService_RemoveARecord(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/custom/example.com", testHandler("dns/removerecord.html", http.MethodDelete, http.StatusOK))

	resp, err := client.DNS.RemoveARecord(context.Background(), "example.com", "1.2.3.4")
	require.NoError(t, err)

	assert.Equal(t, "updated DNS: example.com", resp)
}

func TestDNSService_GetDump(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/admin/dns/dump", testHandler("dns/dump.json", http.MethodGet, http.StatusOK))

	records, err := client.DNS.GetDump(context.Background())
	require.NoError(t, err)

	expected := []Zone{
		{
			Zone: "example1.com",
			Records: []Record{
				{
					Explanation: "Required. Specifies the hostname (and priority) of the machine that handles @example.com mail.",
					Name:        "example1.com",
					Type:        "MX",
					Value:       "10 box.example1.com.",
				},
			},
		},
		{
			Zone: "example2.com",
			Records: []Record{
				{
					Explanation: "Required. Specifies the hostname (and priority) of the machine that handles @example.com mail.",
					Name:        "example2.com",
					Type:        "MX",
					Value:       "10 box.example2.com.",
				},
			},
		},
		{
			Zone: "example3.com",
		},
		{
			Zone: "example4.com",
			Records: []Record{
				{
					Explanation: "Required. Specifies the hostname (and priority) of the machine that handles @example.com mail.",
					Name:        "example4.com",
					Type:        "MX",
					Value:       "10 box.example4.com.",
				},
				{
					Name:  "example4.com",
					Type:  "A",
					Value: "10.0.0.1",
				},
				{
					Name:  "example4.com",
					Type:  "TXT",
					Value: "data",
				},
			},
		},
	}

	assert.Equal(t, expected, records)
}
