package repository

import (
	"context"

	"database/sql"

	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/microservices/user/internal/domain"
	"github.com/Mockird31/OnlineStore/microservices/user/model"
	"github.com/Mockird31/OnlineStore/microservices/user/model/errors"
)

const (
	SignupUserQuery = `
		INSERT INTO "user" (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	CheckUsernameQuery = `
		SELECT 1 
		FROM "user"
		WHERE username = $1
	`
	CheckEmailQuery = `
		SELECT 1 
		FROM "user"
		WHERE email = $1
	`
	GetPasswordHashQuery = `
		SELECT password_hash
		FROM "user"
		WHERE username = $1 OR email = $2
	`
	GetUserIDByUsernameQuery = `
		SELECT id 
		FROM "user"
		WHERE username = $1
	`
)

type userPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) domain.Repository {
	return &userPostgresRepository{
		db: db,
	}
}

func (r *userPostgresRepository) SignupUser(ctx context.Context, username, email, passwordHash string) (*model.User, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	var userID int64
	err := r.db.QueryRowContext(ctx, SignupUserQuery, username, email, passwordHash).Scan(&userID)
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
	logger := loggerPkg.LoggerFromContext(ctx)
	var exist bool
	err := r.db.QueryRowContext(ctx, CheckUsernameQuery, username).Scan(&exist)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		logger.Error("db failed to check username unique")
		return false, errors.NewDatabaseError("wrong database query")
	}
	return exist, nil
}

func (r *userPostgresRepository) CheckEmailUnique(ctx context.Context, email string) (bool, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	var exist bool
	err := r.db.QueryRowContext(ctx, CheckEmailQuery, email).Scan(&exist)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		logger.Error("db failed to check email unique")
		return false, errors.NewDatabaseError("wrong database query")
	}
	return exist, nil
}

func (r *userPostgresRepository) GetPasswordHash(ctx context.Context, username, email string) (string, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	var passwordHash string
	err := r.db.QueryRowContext(ctx, GetPasswordHashQuery, username, email).Scan(&passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("user not exist")
			return "", errors.NewUserNotExistError("user not exist")
		}
		logger.Error("db failed to get password hash")
		return "", errors.NewDatabaseError("wrong database query")
	}
	return passwordHash, nil
}

func (r *userPostgresRepository) GetUserIDByUsername(ctx context.Context, username string) (int64, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	var id int64
	err := r.db.QueryRowContext(ctx, GetUserIDByUsernameQuery, username).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("user not exist")
			return 0, errors.NewUserNotExistError("user not exist")
		}
		logger.Error("db failed to get user id by username")
		return 0, errors.NewDatabaseError("wrong database query")
	}
	return id, nil
}
