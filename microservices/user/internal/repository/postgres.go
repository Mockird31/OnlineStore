package repository

import (
	"context"

	"github.com/Mockird31/OnlineStore/microservices/user/model"
	"github.com/Mockird31/OnlineStore/microservices/user/model/errors"
	"github.com/jackc/pgx/v5/pgxpool"
	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
)

const (
	SignupUserQuery = `
		INSERT INTO user (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	CheckUsernameQuery = `
		SELECT 1 
		FROM user
		WHERE username = $1
	`
	CheckEmailQuery = `
		SELECT 1 
		FROM user
		WHERE email = $1
	`
)

type userPostgresRepository struct {
	db *pgxpool.Pool
}

func (r *userPostgresRepository) SignupUser(ctx context.Context, username, email, passwordHash string) (*model.User, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	var userID int64
	err := r.db.QueryRow(ctx, SignupUserQuery, username, email, passwordHash).Scan(&userID)
	if err != nil {
		logger.Error("db failed to create user")
		return nil, errors.NewDatabaseError("wrong database query")
	}
	user := &model.User{
		Id:       userID,
		Username: username,
		Email:    email,
	}
	return user, nil
}

func (r *userPostgresRepository) CheckUsernameUnique(ctx context.Context, username string) (bool, error) {
	var exist bool
	err := r.db.QueryRow(ctx, CheckUsernameQuery, username).Scan(&exist)
	if err != nil {
		return false, errors.NewDatabaseError("wrong database query")
	}
	return exist, nil
}

func (r *userPostgresRepository) CheckEmailUnique(ctx context.Context, email string) (bool, error) {
	var exist bool
	err := r.db.QueryRow(ctx, CheckEmailQuery, email).Scan(&exist)
	if err != nil {
		return false, errors.NewDatabaseError("wrong database query")
	}
	return exist, nil
}
