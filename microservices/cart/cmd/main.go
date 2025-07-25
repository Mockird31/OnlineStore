package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	userProto "github.com/Mockird31/OnlineStore/gen/user"
	"github.com/Mockird31/OnlineStore/init/postgres"
	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"google.golang.org/grpc"

	"github.com/Mockird31/OnlineStore/config"
	"github.com/Mockird31/OnlineStore/microservices/interceptor"
	"github.com/Mockird31/OnlineStore/microservices/user/internal/delivery"
	"github.com/Mockird31/OnlineStore/microservices/user/internal/repository"
	"github.com/Mockird31/OnlineStore/microservices/user/internal/usecase"
	"go.uber.org/zap"
)

func main() {
	logger, err := loggerPkg.NewZapLogger()
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
		logger.Error("Error loading config:", zap.Error(err))
		return
	}

	port := fmt.Sprintf(":%d", cfg.Services.UserService.Port)
	conn, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("Can't start user service:", zap.Error(err))
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

	postgresPool, err := postgres.ConnectPostgres(cfg.Postgres)
	if err != nil {
		logger.Error("Error connecting to postgres:", zap.Error(err))
		return
	}
	defer func() {
		if err := postgresPool.Close(); err != nil {
			logger.Error("Error closing postgres pool:", zap.Error(err))
		}
	}()

	userRepository := repository.NewUserPostgresRepository(postgresPool)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userService := delivery.NewUserService(userUsecase)
	userProto.RegisterUserServiceServer(server, userService)

	logger.Info("User service started on port %s...", zap.String("port", port))

	err = server.Serve(conn)
	if err != nil {
		logger.Fatal("Error starting user service:", zap.Error(err))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	server.GracefulStop()
	logger.Info("User service stopped")
}
