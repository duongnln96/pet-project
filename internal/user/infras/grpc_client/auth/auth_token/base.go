package auth_token

import (
	"context"

	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/duongnln96/blog-realworld/pkg/config"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var _ port.AuthTokenDomainI = (*authTokenGRPCClient)(nil)

var AuthTokenGrpcClientSet = wire.NewSet(NewGrpcClient)

type authTokenGRPCClient struct {
	conn *grpc.ClientConn
}

func NewGrpcClient(cfg *config.Configs) (port.AuthTokenDomainI, error) {

	clientConfig := cfg.GrpcClientConfigMap.Get("auth_grpc_client")

	ctx, cancel := context.WithTimeout(context.Background(), clientConfig.Timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, clientConfig.Url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                clientConfig.Timeout,
			Timeout:             clientConfig.Timeout,
			PermitWithoutStream: false,
		}),
	)
	if err != nil {
		return nil, err
	}

	return &authTokenGRPCClient{
		conn: conn,
	}, nil
}
