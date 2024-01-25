package auth_token

import (
	"context"
	"log"
	"sync"

	"github.com/duongnln96/blog-realworld/internal/user/core/port"
	"github.com/duongnln96/blog-realworld/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	authTokenGen "github.com/duongnln96/blog-realworld/gen/go/auth/v1"
)

var (
	authTokenDomainInstance port.AuthTokenDomainI
	authTokenGRPCClientOnce = sync.Once{}
)

var _ port.AuthTokenDomainI = (*authTokenDomainClient)(nil)

type authTokenDomainClient struct {
	conn   *grpc.ClientConn
	domain authTokenGen.AuthTokenServiceClient
}

func NewAuthTokenDomainClient(cfg *config.Configs) port.AuthTokenDomainI {

	if authTokenDomainInstance != nil {
		return authTokenDomainInstance
	}

	authTokenGRPCClientOnce.Do(func() {
		clientConfig := cfg.GrpcClientConfigMap.Get("auth_grpc_client")

		ctx, cancel := context.WithTimeout(context.Background(), clientConfig.Timeout)
		defer cancel()

		conn, err := grpc.DialContext(ctx, clientConfig.Url,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:    clientConfig.Timeout,
				Timeout: clientConfig.Timeout,
			}),
		)
		if err != nil {
			log.Fatalf("grpc.DialContext %s - %s", clientConfig.Url, err.Error())
			return
		}

		authTokenDomainInstance = &authTokenDomainClient{
			conn:   conn,
			domain: authTokenGen.NewAuthTokenServiceClient(conn),
		}
	})

	return authTokenDomainInstance
}
