package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/Mockird31/OnlineStore/config"
	"github.com/Mockird31/OnlineStore/init/redis"
	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/microservices/interceptor"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/Mockird31/OnlineStore/microservices/auth/internal/delivery"
	"github.com/Mockird31/OnlineStore/microservices/auth/internal/repository"
	"github.com/Mockird31/OnlineStore/microservices/auth/internal/usecase"

	authProto "github.com/Mockird31/OnlineStore/gen/auth"
)

func main() {
	logger, err := logger.NewZapLogger()
	if err != nil {
		logger.Error("Error creating logger:", zap.Error(err))
		return
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Error("Error syncing logger:", zap.Error(err))
		}
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Fail to load config:", zap.Error(err))
	}

	port := fmt.Sprintf(":%d", cfg.Services.AuthService.Port)
	conn, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("Can't start auth service:", zap.Error(err))
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Error("Error closing connection:", zap.Error(err))
		}
	}()

	accessInterceptor := interceptor.NewAccessInterceptor(logger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(accessInterceptor.UnaryServerInterceptor()),
	)

	redis := redis.NewRedisPool(cfg.Redis)
	logger.Info("success connect to redis")
	defer func() {
		if err := redis.Close(); err != nil {
			logger.Error("Error closing redis:", zap.Error(err))
		}
	}()

	authRepository := repository.NewAuthRepository(redis)
	authUsecase := usecase.NewAuthUsecase(authRepository)
	authService := delivery.NewAuthService(authUsecase)
	authProto.RegisterAuthServiceServer(server, authService)

	logger.Info("Auth service started on port: %s", cfg.Services.AuthService.Port)

	err = server.Serve(conn)
	if err != nil {
		logger.Fatal("Error while starting auth service:", zap.Error(err))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	server.GracefulStop()
	logger.Info("Auth service graceful stopped")
}
