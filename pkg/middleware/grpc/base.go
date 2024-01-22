package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type LoggingInterceptor interface {
	LoggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
}
