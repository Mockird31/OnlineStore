package interceptor

import (
	"context"
	"time"

	"github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AccessInterceptor struct {
	logger *zap.SugaredLogger
}

func NewAccessInterceptor(logger *zap.SugaredLogger) *AccessInterceptor {
	return &AccessInterceptor{logger: logger}
}

func (i *AccessInterceptor) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctxLogger := i.logger.With(zap.String("method", info.FullMethod))
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			requestID := md.Get("request_id")
			if len(requestID) > 0 {
				ctxLogger = ctxLogger.With(zap.String("request_id", requestID[0]))
			}
		}

		newCtx := logger.LoggerToContext(ctx, ctxLogger)

		start := time.Now()

		ctxLogger.Infow("gRPC request received",
			"method", info.FullMethod,
			"request", req,
		)

		resp, err = handler(newCtx, req)

		duration := time.Since(start)

		if err != nil {
			st, _ := status.FromError(err)
			ctxLogger.Errorw("gRPC request failed",
				"method", info.FullMethod,
				"code", st.Code(),
				"error", err,
				"duration", duration,
			)
		} else {
			ctxLogger.Infow("gRPC request completed",
				"method", info.FullMethod,
				"duration", duration,
			)
		}
		return resp, err
	}
}
