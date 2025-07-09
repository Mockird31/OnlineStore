package middleware

import (
	"context"
	"net/http"

	loggerHelper "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type RequestIDKey struct{}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), RequestIDKey{}, requestID)
		logger := loggerHelper.LoggerFromContext(ctx).With(zap.String("request_id", requestID))
		ctx = loggerHelper.LoggerToContext(ctx, logger)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
