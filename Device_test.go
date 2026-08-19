package onvif

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// vivotekGetServicesResponse is based on a real GetServices response from a
// Vivotek SD9161-H, which exposes ver10 media and ver20 media at genuinely
// distinct XAddrs.
const vivotekGetServicesResponse = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope">
  <SOAP-ENV:Body>
    <tds:GetServicesResponse xmlns:tds="http://www.onvif.org/ver10/device/wsdl">
      <tds:Service>
        <tds:Namespace>http://www.onvif.org/ver10/device/wsdl</tds:Namespace>
        <tds:XAddr>http://192.168.1.100/onvif/device_service</tds:XAddr>
        <tds:Version><tt:Major>2</tt:Major><tt:Minor>40</tt:Minor></tds:Version>
      </tds:Service>
      <tds:Service>
        <tds:Namespace>http://www.onvif.org/ver10/media/wsdl</tds:Namespace>
        <tds:XAddr>http://192.168.1.100/onvif/media_service</tds:XAddr>
        <tds:Version><tt:Major>2</tt:Major><tt:Minor>40</tt:Minor></tds:Version>
      </tds:Service>
      <tds:Service>
        <tds:Namespace>http://www.onvif.org/ver20/media/wsdl</tds:Namespace>
        <tds:XAddr>http://192.168.1.100/onvif/media2_service</tds:XAddr>
        <tds:Version><tt:Major>2</tt:Major><tt:Minor>40</tt:Minor></tds:Version>
      </tds:Service>
      <tds:Service>
        <tds:Namespace>http://www.onvif.org/ver20/ptz/wsdl</tds:Namespace>
        <tds:XAddr>http://192.168.1.100/onvif/ptz_service</tds:XAddr>
        <tds:Version><tt:Major>2</tt:Major><tt:Minor>40</tt:Minor></tds:Version>
      </tds:Service>
    </tds:GetServicesResponse>
  </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

// getCapabilitiesOnlyResponse represents an older camera that only supports
// GetCapabilities, which does not report a distinct Media2 endpoint.
const getCapabilitiesOnlyResponse = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope">
  <SOAP-ENV:Body>
    <tds:GetCapabilitiesResponse xmlns:tds="http://www.onvif.org/ver10/device/wsdl">
      <tds:Capabilities>
        <tt:Media xmlns:tt="http://www.onvif.org/ver10/schema">
          <tt:XAddr>http://192.168.1.100/onvif/media_service</tt:XAddr>
        </tt:Media>
        <tt:PTZ xmlns:tt="http://www.onvif.org/ver10/schema">
          <tt:XAddr>http://192.168.1.100/onvif/ptz_service</tt:XAddr>
        </tt:PTZ>
      </tds:Capabilities>
    </tds:GetCapabilitiesResponse>
  </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

func newTestDevice(t *testing.T) *Device {
	t.Helper()
	dev := &Device{
		params:    DeviceParams{Xaddr: "192.168.1.100"},
		endpoints: make(map[string]string),
	}
	return dev
}

func TestDevice_GetServices_DistinctMedia2Endpoint(t *testing.T) {
	dev := newTestDevice(t)
	resp := &http.Response{Body: io.NopCloser(strings.NewReader(vivotekGetServicesResponse))}

	dev.getServicesFromGetServices(resp)
	dev.fillMediaFallback()

	assert.Equal(t, "http://192.168.1.100/onvif/media_service", dev.endpoints["media"])
	assert.Equal(t, "http://192.168.1.100/onvif/media2_service", dev.endpoints["media2"])
}

// A GetServices response that omits services GetCapabilities did report must not
// drop them: GetServices is an overlay, never a replacement.
func TestDevice_GetServices_IsOverlayNotReplacement(t *testing.T) {
	dev := newTestDevice(t)
	capsResp := &http.Response{Body: io.NopCloser(strings.NewReader(getCapabilitiesOnlyResponse))}
	dev.getSupportedServices(capsResp)

	const partialGetServices = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://www.w3.org/2003/05/soap-envelope">
  <SOAP-ENV:Body>
    <tds:GetServicesResponse xmlns:tds="http://www.onvif.org/ver10/device/wsdl">
      <tds:Service>
        <tds:Namespace>http://www.onvif.org/ver20/media/wsdl</tds:Namespace>
        <tds:XAddr>http://192.168.1.100/onvif/media2_service</tds:XAddr>
      </tds:Service>
    </tds:GetServicesResponse>
  </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`
	dev.getServicesFromGetServices(&http.Response{Body: io.NopCloser(strings.NewReader(partialGetServices))})
	dev.fillMediaFallback()

	// The distinct Media2 XAddr is picked up...
	assert.Equal(t, "http://192.168.1.100/onvif/media2_service", dev.endpoints["media2"])
	// ...without discarding endpoints only GetCapabilities reported.
	assert.Equal(t, "http://192.168.1.100/onvif/media_service", dev.endpoints["media"])
	assert.Equal(t, "http://192.168.1.100/onvif/ptz_service", dev.endpoints["ptz"])
}

func TestDevice_GetCapabilities_Media2FallsBackToMedia(t *testing.T) {
	dev := newTestDevice(t)
	resp := &http.Response{Body: io.NopCloser(strings.NewReader(getCapabilitiesOnlyResponse))}

	dev.getSupportedServices(resp)
	dev.fillMediaFallback()

	assert.Equal(t, "http://192.168.1.100/onvif/media_service", dev.endpoints["media"])
	assert.Equal(t, "http://192.168.1.100/onvif/media_service", dev.endpoints["media2"])
}

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
