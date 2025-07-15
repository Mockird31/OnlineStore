package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func setupTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, context.Context) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	logger, err := config.Build()
	if err != nil {
		logger = zap.NewNop()
	}

	ctx := loggerPkg.LoggerToContext(context.Background(), logger.Sugar())

	return db, mock, ctx
}

func TestSignupUser(t *testing.T) {
	db, mock, ctx := setupTest(t)
	defer db.Close()

	repo := NewUserPostgresRepository(db)
	username := "username"
	email := "email"
	passwordHash := "passwordHash"
	testID := int64(1)

	mock.ExpectQuery("INSERT INTO \"user\"").WithArgs(username, email, passwordHash).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(testID))

	user, err := repo.SignupUser(ctx, username, email, passwordHash)

	assert.NoError(t, err)
	assert.Equal(t, testID, user.Id)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func TestCheckUsernameUnique(t *testing.T) {
	db, mock, ctx := setupTest(t)
	defer db.Close()

	repo := NewUserPostgresRepository(db)
	username := "username"

	mock.ExpectQuery("SELECT 1").WithArgs(username).WillReturnError(sql.ErrNoRows)

	isUsernameExist, err := repo.CheckUsernameUnique(ctx, username)
	assert.NoError(t, err)
	assert.Equal(t, isUsernameExist, false)
}

func TestCheckUsernameUniqueFalse(t *testing.T) {
	db, mock, ctx := setupTest(t)
	defer db.Close()

	repo := NewUserPostgresRepository(db)
	username := "username"

	mock.ExpectQuery("SELECT 1").WithArgs(username).WillReturnRows(
		sqlmock.NewRows([]string{"1"}).AddRow(1))

	isUsernameExist, err := repo.CheckUsernameUnique(ctx, username)
	assert.NoError(t, err)
	assert.Equal(t, isUsernameExist, true)
}

func TestCheckEmailUnique(t *testing.T) {
	db, mock, ctx := setupTest(t)
	defer db.Close()

	repo := NewUserPostgresRepository(db)
	email := "email"

	mock.ExpectQuery("SELECT 1").WithArgs(email).WillReturnError(sql.ErrNoRows)

	isUsernameExist, err := repo.CheckEmailUnique(ctx, email)
	assert.NoError(t, err)
	assert.Equal(t, isUsernameExist, false)
}

func TestCheckEmailUniqueFalse(t *testing.T) {
	db, mock, ctx := setupTest(t)
	defer db.Close()

	repo := NewUserPostgresRepository(db)
	email := "email"

	mock.ExpectQuery("SELECT 1").WithArgs(email).WillReturnRows(
		sqlmock.NewRows([]string{"1"}).AddRow(1))

	isUsernameExist, err := repo.CheckEmailUnique(ctx, email)
	assert.NoError(t, err)
	assert.Equal(t, isUsernameExist, true)
}

func TestGetPasswordHashByUsername(t *testing.T) {
	db, mock, ctx := setupTest(t)
	defer db.Close()

	repo := NewUserPostgresRepository(db)
	username := "username"

	mock.ExpectQuery("SELECT password_hash").WithArgs(username, "").WillReturnRows(sqlmock.NewRows([]string{"hash"}).AddRow("hash"))

	passwordHash, err := repo.GetPasswordHash(ctx, username, "")
	assert.NoError(t, err)
	assert.Equal(t, passwordHash, "hash")
}

func TestGetPasswordHashByUsernameErrorNoHash(t *testing.T) {
	db, mock, ctx := setupTest(t)
	defer db.Close()

	repo := NewUserPostgresRepository(db)
	username := "username"

	mock.ExpectQuery("SELECT password_hash").WithArgs(username, "").WillReturnError(sql.ErrNoRows)

	passwordHash, err := repo.GetPasswordHash(ctx, username, "")
	assert.Error(t, err)
	assert.Equal(t, passwordHash, "")
}

func TestGetPasswordHashByEmail(t *testing.T) {
	db, mock, ctx := setupTest(t)
	defer db.Close()

	repo := NewUserPostgresRepository(db)
	email := "email"

	mock.ExpectQuery("SELECT password_hash").WithArgs("", email).WillReturnRows(sqlmock.NewRows([]string{"hash"}).AddRow("hash"))

	passwordHash, err := repo.GetPasswordHash(ctx, "", email)
	assert.NoError(t, err)
	assert.Equal(t, passwordHash, "hash")
}

func TestGetPasswordHashByEmailErrorNoHash(t *testing.T) {
	db, mock, ctx := setupTest(t)
	defer db.Close()

	repo := NewUserPostgresRepository(db)
	email := "email"

	mock.ExpectQuery("SELECT password_hash").WithArgs("", email).WillReturnError(sql.ErrNoRows)

	passwordHash, err := repo.GetPasswordHash(ctx, "", email)
	assert.Error(t, err)
	assert.Equal(t, passwordHash, "")
}

func TestGetUserIDByUsername(t *testing.T) {
	db, mock, ctx := setupTest(t)
	defer db.Close()

	repo := NewUserPostgresRepository(db)

	username := "username"
	userIDTest := int64(1)

	mock.ExpectQuery("SELECT id").WithArgs(username).WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))

	userID, err := repo.GetUserIDByUsername(ctx, username)
	assert.NoError(t, err)
	assert.Equal(t, userIDTest, userID)
}

func TestGetUserIDByUsernameError(t *testing.T) {
	db, mock, ctx := setupTest(t)
	defer db.Close()

	repo := NewUserPostgresRepository(db)

	username := "username"

	mock.ExpectQuery("SELECT id").WithArgs(username).WillReturnError(sql.ErrNoRows)

	_, err := repo.GetUserIDByUsername(ctx, username)
	assert.Error(t, err)
}