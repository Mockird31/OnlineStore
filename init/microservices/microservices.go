package microservices

import (
	"context"
	"fmt"

	"github.com/Mockird31/OnlineStore/config"
	"github.com/Mockird31/OnlineStore/internal/middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type Clients struct {
	AuthClient *grpc.ClientConn
	UserClient *grpc.ClientConn
}

func requestIdUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	requestId := ctx.Value(middleware.RequestIDKey{}).(string)
	md := metadata.New(map[string]string{
		"request_id": requestId,
	})
	return invoker(metadata.NewOutgoingContext(ctx, md), method, req, reply, cc, opts...)
}

func InitMicroservices(cfg *config.Services, logger *zap.SugaredLogger) (*Clients, error) {
	authAddress := fmt.Sprintf("%s:%d", cfg.AuthService.Host, cfg.AuthService.Port)
	authClient, err := grpc.NewClient(authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(requestIdUnaryClientInterceptor))
	if err != nil {
		return nil, err
	}

	userAddress := fmt.Sprintf("%s:%d", cfg.UserService.Host, cfg.UserService.Port)
	userClient, err := grpc.NewClient(userAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(requestIdUnaryClientInterceptor))
	if err != nil {
		return nil, err
	}
	return &Clients{
		AuthClient: authClient,
		UserClient: userClient,
	}, nil
}
