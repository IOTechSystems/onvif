package wsdiscovery

import (
	"context"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const probeMatch = `<Envelope><Body><ProbeMatches><ProbeMatch><XAddrs>http://10.0.0.5/onvif/device_service</XAddrs></ProbeMatch></ProbeMatches></Body></Envelope>`

// fakeCamera answers one Probe with a ProbeMatch.
func fakeCamera(t *testing.T) *net.UDPAddr {
	t.Helper()
	camera, err := net.ListenPacket("udp4", "127.0.0.1:0")
	require.NoError(t, err)
	t.Cleanup(func() { _ = camera.Close() })
	go func() {
		buf := make([]byte, bufSize)
		n, from, err := camera.ReadFrom(buf)
		if err != nil || !strings.Contains(string(buf[:n]), "discovery/Probe") {
			return
		}
		_, _ = camera.WriteTo([]byte(probeMatch), from)
	}()
	return camera.LocalAddr().(*net.UDPAddr)
}

func TestSendUnicastProbe(t *testing.T) {
	camera := fakeCamera(t)
	silent := fakeCamera(t)
	silent.Port++ // nothing listens here

	responses, err := SendUnicastProbe(context.Background(), nil, []*net.UDPAddr{camera, silent},
		200*time.Millisecond, nil, []string{"dn:NetworkVideoTransmitter"}, nil)
	require.NoError(t, err)
	require.Len(t, responses, 1)
	assert.Equal(t, camera.String(), responses[0].Addr.String())
	assert.Equal(t, probeMatch, responses[0].Body)
}

func TestSendUnicastProbeCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	start := time.Now()
	_, err := SendUnicastProbe(ctx, nil, []*net.UDPAddr{fakeCamera(t)}, 10*time.Second, nil, nil, nil)
	assert.ErrorIs(t, err, context.Canceled)
	assert.Less(t, time.Since(start), time.Second)
}
