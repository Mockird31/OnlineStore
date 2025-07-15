package repository

import (
	"context"
	"testing"
	"time"

	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock/v3"
	"go.uber.org/zap"
)

func setupTest() (*redis.Pool, *redigomock.Conn, context.Context) {
	conn := redigomock.NewConn()

	pool := &redis.Pool{
		MaxIdle:     3,
		MaxActive:   10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return conn, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			return nil
		},
	}

	logger := zap.NewNop().Sugar()
	ctx := loggerPkg.LoggerToContext(context.Background(), logger)

	return pool, conn, ctx
}

func TestCreateSession(t *testing.T) {
	redis, conn, ctx := setupTest()
	defer redis.Close()
	data := []byte("some data")
	repo := NewAuthRepository(redis)

	conn.Command("SETEX", redigomock.NewAnyData(), 86400, data).Expect("OK")

	sessionID, err := repo.CreateSession(ctx, data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if sessionID == "" {
		t.Fatalf("expected non-empty session ID, got empty string")
	}
}

func TestGetUserBySessionID(t *testing.T) {
	redis, conn, ctx := setupTest()
	defer redis.Close()
	repo := NewAuthRepository(redis)

	sessionID := "something"
	resUser := []byte("something data")
	conn.Command("GET", sessionID).Expect(resUser)

	user, err := repo.GetUserBySessionID(ctx, sessionID)
	if err != nil {
		t.Fatalf("failed to get user %v", err)
	}

	if string(user) != string(resUser) {
		t.Fatalf("different users")
	}
}

func TestDeleteSession(t *testing.T) {
	redis, conn, ctx := setupTest()
	defer redis.Close()
	repo := NewAuthRepository(redis)

	sessionID := "something"
	conn.Command("DEL", sessionID).Expect("OK")

	err := repo.DeleteSession(ctx, sessionID)
	if err != nil {
		t.Fatalf("failed to delete session")
	}
}