package transport

import (
	"context"

	"github.com/faireal/trojan-go/tunnel"
)

const Name = "TRANSPORT"

type Tunnel struct{}

func (*Tunnel) Name() string {
	return Name
}

func (*Tunnel) NewClient(ctx context.Context, client tunnel.Client) (tunnel.Client, error) {
	return NewClient(ctx, client)
}

func (*Tunnel) NewServer(ctx context.Context, server tunnel.Server) (tunnel.Server, error) {
	return NewServer(ctx, server)
}

func init() {
	tunnel.RegisterTunnel(Name, &Tunnel{})
}
