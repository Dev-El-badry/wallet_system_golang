package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardForHeader          = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		// log.Printf("md: %+v\n", md)
		if userAgent := md.Get(grpcGatewayUserAgentHeader); len(userAgent) > 0 {
			mtdt.UserAgent = userAgent[0]
		}

		if userAgent := md.Get(userAgentHeader); len(userAgent) > 0 {
			mtdt.UserAgent = userAgent[0]
		}

		if clientIp := md.Get(xForwardForHeader); len(clientIp) > 0 {
			mtdt.ClientIP = clientIp[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		mtdt.UserAgent = p.Addr.String()
	}

	return mtdt
}
