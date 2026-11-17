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
func SendUnicastProbe(ctx context.Context, conn net.PacketConn, targets []*net.UDPAddr, timeout time.Duration, scopes, types []string, namespaces map[string]string) ([]ProbeResponse, error) {
	probeSOAP := BuildProbeMessage(uuid.NewString(), scopes, types, namespaces)
	return SendUDPUnicast(ctx, conn, probeSOAP.String(), targets, timeout)
}

// SendUDPUnicast writes msg to each target and reads replies until timeout or ctx
// is done. A nil conn listens on an ephemeral udp4 port; a non-nil one is not closed.
func SendUDPUnicast(ctx context.Context, conn net.PacketConn, msg string, targets []*net.UDPAddr, timeout time.Duration) ([]ProbeResponse, error) {
	if len(targets) == 0 {
		return nil, errors.New("no unicast targets given")
	}
	if conn == nil {
		var lc net.ListenConfig
		c, err := lc.ListenPacket(ctx, "udp4", "0.0.0.0:0")
		if err != nil {
			return nil, err
		}
		defer func() { _ = c.Close() }()
		conn = c
	}

	data := []byte(msg)
	var sendErr error
	sent := 0
	for _, target := range targets {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		if _, err := conn.WriteTo(data, target); err != nil {
			sendErr = fmt.Errorf("failed to write ws-discovery probe to %s: %w", target, err)
			continue
		}
		sent++
	}
	if sent == 0 {
		return nil, sendErr
	}

	if err := conn.SetReadDeadline(time.Now().Add(timeout)); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}
	fired := make(chan struct{})
	stop := context.AfterFunc(ctx, func() {
		_ = conn.SetReadDeadline(time.Now())
		close(fired)
	})
	// net.PacketConn can't report its previous deadline, so clear it rather than restore it.
	defer func() {
		if !stop() {
			<-fired
		}
		_ = conn.SetReadDeadline(time.Time{})
	}()

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
	if err := ctx.Err(); err != nil {
		return responses, err
	}
	return responses, nil
}
