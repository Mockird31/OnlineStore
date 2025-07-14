package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Mockird31/OnlineStore/config"
	"github.com/Mockird31/OnlineStore/init/microservices"
	"github.com/Mockird31/OnlineStore/init/postgres"
	"github.com/Mockird31/OnlineStore/internal/middleware"
	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	authProto "github.com/Mockird31/OnlineStore/gen/auth"
	userProto "github.com/Mockird31/OnlineStore/gen/user"

	userHttp "github.com/Mockird31/OnlineStore/internal/pkg/user/delivery/http"
	userUsecase "github.com/Mockird31/OnlineStore/internal/pkg/user/usecase"
)

func main() {
	logger, err := logger.NewZapLogger()
	if err != nil {
		logger.Error("Error creating logger:", zap.Error(err))
		return
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Error loading config:", zap.Error(err))
		return
	}

	postgresConn, err := postgres.ConnectPostgres(cfg.Postgres)
	if err != nil {
		logger.Error("Error connecting to Postgres:", zap.Error(err))
		return
	}
	defer func() {
		if err := postgresConn.Close(); err != nil {
			logger.Error("Error closing Postgres:", zap.Error(err))
		}
	}()

	err = postgres.RunMigrations(cfg.Postgres)
	if err != nil {
		logger.Error("Error running migrations:", zap.Error(err))
		return
	}

	r := mux.NewRouter()
	logger.Info("Server starting on port %s...", zap.String("port", fmt.Sprintf(":%d", cfg.Port)))

	clients, err := microservices.InitMicroservices(&cfg.Services, logger)
	if err != nil {
		logger.Error("Error initializing gRPC clients:", zap.Error(err))
		return
	}

	authClient := authProto.NewAuthServiceClient(clients.AuthClient)
	userClient := userProto.NewUserServiceClient(clients.UserClient)

	r.Use(middleware.LoggerMiddleware(logger))
	r.Use(middleware.RequestID)
	r.Use(middleware.IsAuth(authClient))

	userHandler := userHttp.NewUserHandler(userUsecase.NewUserUsecase(userClient, authClient))

	r.HandleFunc("/api/v1/auth/signup", userHandler.SignupUser).Methods("POST")
	r.HandleFunc("/api/v1/auth/login", userHandler.LoginUser).Methods("POST")
	r.HandleFunc("/api/v1/auth/logout", userHandler.LogoutUser).Methods("POST")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}

	err = srv.ListenAndServe()
	if err != nil {
		logger.Error("Error starting server:", zap.Error(err))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		logger.Error("Error shutting down server:", zap.Error(err))
		os.Exit(1)
	}
	logger.Info("Composer server stopped")
}
