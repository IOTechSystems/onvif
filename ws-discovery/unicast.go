package wsdiscovery

// SPDX-License-Identifier: Apache-2.0

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/google/uuid"
)

// ProbeResponse is a raw ws-discovery reply and the address it came from.
type ProbeResponse struct {
	Addr *net.UDPAddr
	Body string
}

// SendUnicastProbe sends a Probe to each target and collects the replies. It
// reaches devices multicast can't, such as those across a routed VPN link.
func SendUnicastProbe(ctx context.Context, targets []*net.UDPAddr, timeout time.Duration, scopes, types []string, namespaces map[string]string) ([]ProbeResponse, error) {
	probeSOAP := BuildProbeMessage(uuid.NewString(), scopes, types, namespaces)
	return SendUDPUnicast(ctx, probeSOAP.String(), targets, timeout)
}

// SendUDPUnicast writes msg to each target from an ephemeral udp4 port and reads
// replies until timeout or ctx is done. Replies collected before a failed write
// or a cancelled ctx are returned alongside the error.
func SendUDPUnicast(ctx context.Context, msg string, targets []*net.UDPAddr, timeout time.Duration) ([]ProbeResponse, error) {
	if len(targets) == 0 {
		return nil, errors.New("no unicast targets given")
	}
	var lc net.ListenConfig
	conn, err := lc.ListenPacket(ctx, "udp4", "0.0.0.0:0")
	if err != nil {
		return nil, err
	}
	defer func() { _ = conn.Close() }()

	data := []byte(msg)
	var sendErrs []error
	for _, target := range targets {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		if _, err := conn.WriteTo(data, target); err != nil {
			sendErrs = append(sendErrs, fmt.Errorf("failed to write ws-discovery probe to %s: %w", target, err))
		}
	}
	if len(sendErrs) == len(targets) {
		return nil, errors.Join(sendErrs...)
	}

	if err := conn.SetReadDeadline(time.Now().Add(timeout)); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}
	stop := context.AfterFunc(ctx, func() { _ = conn.SetReadDeadline(time.Now()) })
	defer stop()

	var responses []ProbeResponse
	buf := make([]byte, bufSize)
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			if !errors.Is(err, os.ErrDeadlineExceeded) {
				return nil, fmt.Errorf("unexpected error occurred while reading ws-discovery responses: %w", err)
			}
			break
		}
		if udpAddr, ok := addr.(*net.UDPAddr); ok {
			responses = append(responses, ProbeResponse{Addr: udpAddr, Body: string(buf[:n])})
		}
	}
	return responses, errors.Join(append(sendErrs, ctx.Err())...)
}
