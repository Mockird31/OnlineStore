package repository

import (
	"context"
	"time"

	domain "github.com/Mockird31/OnlineStore/microservices/auth/internal/domain"
	errors "github.com/Mockird31/OnlineStore/microservices/auth/model/errors"
	"github.com/gomodule/redigo/redis"

	"crypto/rand"
	"encoding/hex"
)

const (
	EXPIRATION = time.Hour * 24
)

type authRepository struct {
	redis *redis.Pool
}

func NewAuthRepository(redis *redis.Pool) domain.Repository {
	return &authRepository{redis: redis}
}

func generateSessionID() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", errors.NewGenerateSessionError("failed to generate session id")
	}
	return hex.EncodeToString(bytes), nil
}

func (r *authRepository) CreateSession(ctx context.Context, user []byte) (string, error) {
	conn := r.redis.Get()
	defer conn.Close()

	sessionID, err := generateSessionID()
	if err != nil {
		return "", err
	}
	_, err = redis.DoContext(conn, ctx, "SETEX", sessionID, int(EXPIRATION.Seconds()), user)
	if err != nil {
		return "", errors.NewSetSessionIDError("failed to set session")
	}
	return sessionID, nil
}

func (r *authRepository) GetUserBySessionID(ctx context.Context, sessionID string) ([]byte, error) {
	conn := r.redis.Get()
	defer conn.Close()

	userBytes, err := redis.DoContext(conn, ctx, "GET", sessionID)
	if err != nil {
		return nil, errors.NewGetSessionError("failed to get user id")
	}

	if userBytes == nil {
		return nil, errors.NewFindSessionError("failed to find user id by session id")
	}

	data, ok := userBytes.([]byte)
	if !ok {
		return nil, errors.NewFailToParseRedisIntError("failed to parse redis value")
	}

	return data, nil
}

func (r *authRepository) DeleteSession(ctx context.Context, sessionID string) error {
	conn := r.redis.Get()
	defer conn.Close()

	_, err := redis.DoContext(conn, ctx, "DEL", sessionID)
	if err != nil {
		return errors.NewDeleteSessionError("failed to delete session")
	}

	return nil
}
