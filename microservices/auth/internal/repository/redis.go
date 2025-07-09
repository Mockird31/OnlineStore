package repository

import (
	"context"
	"time"

	domain "github.com/Mockird31/OnlineStore/microservices/auth/internal/domain"
	errors "github.com/Mockird31/OnlineStore/microservices/auth/models/errors"
	"github.com/redis/go-redis/v9"

	"crypto/rand"
	"encoding/hex"
)

const (
	EXPIRATION = time.Hour * 24
)

type authRepository struct {
	redis *redis.Client
}

func NewAuthRepository(redis *redis.Client) domain.Repository {
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

func (r *authRepository) CreateSession(ctx context.Context, userID int64) (string, error) {
	sessionID, err := generateSessionID()
	if err != nil {
		return "", err
	}
	err = r.redis.Set(ctx, sessionID, userID, EXPIRATION).Err()
	if err != nil {
		return "", errors.NewSetSessionIDError("failed to set session id(%s) to user(%d)", sessionID, userID)
	}
	return sessionID, nil
}

func (r *authRepository) GetUserIDBySessionID(ctx context.Context, sessionID string) (int64, error) {
	cmd := r.redis.Get(ctx, sessionID)
	if err := cmd.Err(); err != nil {
		if err == redis.Nil {
			return 0, errors.NewFindSessionError("failed to find user by session id(%s)", sessionID)
		}
		return 0, errors.NewGetSessionError("failed to get user by session id(%s)", sessionID)
	}
	userID, err := cmd.Int64()
	if err != nil {
		return 0, errors.NewFailToParseRedisIntError("failed to parse redis int in session id(%s)", sessionID)
	}
	return userID, nil
}

func (r *authRepository) DeleteSession(ctx context.Context, sessionID string) error {
	err := r.redis.Del(ctx, sessionID).Err()
	if err != nil {
		return errors.NewDeleteSessionError("failed to delete value with session id(%s)", sessionID)
	}
	return nil
}
