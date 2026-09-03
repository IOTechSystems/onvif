package onvif

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDevice_SetDeviceInfoFromScopes(t *testing.T) {
	const (
		name     = "DeviceName"
		hardware = "M9000"
	)
	scopes := []string{
		"onvif://www.onvif.org/Profile/Streaming",
		"onvif://www.onvif.org/SomethingElse/value",
		"onvif://www.onvif.org/name/" + name,
		"onvif://www.onvif.org/hardware/" + hardware,
	}
	device := Device{}
	device.SetDeviceInfoFromScopes(scopes)
	assert.Equal(t, device.info.Name, name)
	assert.Equal(t, device.info.Model, hardware)
}

func TestDevice_now(t *testing.T) {
	fixed := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)

	dev := Device{params: DeviceParams{Now: func() time.Time { return fixed }}}
	assert.Equal(t, fixed, dev.now())

	dev = Device{}
	assert.WithinDuration(t, time.Now(), dev.now(), time.Second)
}

func TestDevice_UsesNowForWSSecurityCreated(t *testing.T) {
	// Camera clock is skewed: a fixed non-UTC "now" must render as UTC in the WS-Security Created.
	fixed := time.Date(2020, 1, 2, 3, 4, 5, 600000000, time.FixedZone("CET", 3600))

	srv := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(b), ">2020-01-02T02:04:05.6Z</Created>")
	}))
	defer srv.Close()

	dev := Device{params: DeviceParams{
		Username:   "user",
		Password:   "pass",
		AuthMode:   UsernameTokenAuth,
		HttpClient: srv.Client(),
		Now:        func() time.Time { return fixed },
	}}

	// Every sender that builds a WS-UsernameToken header must go through dev.now().
	resp, err := dev.SendSoap(srv.URL, "<tds:GetSystemDateAndTime/>")
	require.NoError(t, err)
	resp.Body.Close()

	resp, err = dev.SendGetSnapshotRequest(srv.URL)
	require.NoError(t, err)
	resp.Body.Close()
}
