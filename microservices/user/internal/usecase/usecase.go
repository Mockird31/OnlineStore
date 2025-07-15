package usecase

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"strings"

	loggerPkg "github.com/Mockird31/OnlineStore/internal/pkg/helpers/logger"
	"github.com/Mockird31/OnlineStore/microservices/user/internal/domain"
	"github.com/Mockird31/OnlineStore/microservices/user/model"
	"github.com/Mockird31/OnlineStore/microservices/user/model/errors"
	"golang.org/x/crypto/argon2"
)

type userUsecase struct {
	userPostgresRepository domain.Repository
}

func NewUserUsecase(userPostgresRepository domain.Repository) domain.Usecase {
	return &userUsecase{userPostgresRepository: userPostgresRepository}
}

func CreateSalt() []byte {
	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		return nil
	}
	return salt
}

func CheckPasswordHash(encodedHash string, password string) bool {
	decodedHash, err := base64.StdEncoding.DecodeString(encodedHash)
	if err != nil {
		return false
	}
	salt := decodedHash[:8]
	userPassHash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return bytes.Equal(userPassHash, decodedHash[8:])
}

func HashPassword(salt []byte, password string) string {
	hashedPass := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	combined := append(salt, hashedPass...)
	return base64.StdEncoding.EncodeToString(combined)
}

func (u *userUsecase) SignupUser(ctx context.Context, regData *model.RegisterData) (*model.User, error) {
	regData.Username = strings.ToLower(regData.Username)
	logger := loggerPkg.LoggerFromContext(ctx)
	if regData.Password != regData.ConfirmPassword {
		logger.Error("different passwords")
		return nil, errors.NewWrongPasswordError("wrong password was entered")
	}
	isUsernameExist, err := u.userPostgresRepository.CheckUsernameUnique(ctx, regData.Username)
	if err != nil {
		return nil, err
	}
	if isUsernameExist {
		logger.Error("username not unique")
		return nil, errors.NewNotUniqueError("username %s not unique", regData.Username)
	}
	isEmailExist, err := u.userPostgresRepository.CheckEmailUnique(ctx, regData.Email)
	if err != nil {
		return nil, err
	}
	if isEmailExist {
		logger.Error("email not unique")
		return nil, errors.NewNotUniqueError("email %s not unique", regData.Email)
	}
	salt := CreateSalt()
	if salt == nil {
		return nil, errors.NewCreateSaltError("failed to create salt")
	}
	passwordHash := HashPassword(salt, regData.Password)
	user, err := u.userPostgresRepository.SignupUser(ctx, regData.Username, regData.Email, passwordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) LoginUser(ctx context.Context, logData *model.LoginData) (*model.User, error) {
	logger := loggerPkg.LoggerFromContext(ctx)
	var isUsernameExist, isEmailExist bool
	var err error
	if logData.Username != "" {
		logData.Username = strings.ToLower(logData.Username)
		isUsernameExist, err = u.userPostgresRepository.CheckUsernameUnique(ctx, logData.Username)
		if err != nil {
			return nil, err
		}
	}
	if logData.Email != "" {
		isEmailExist, err = u.userPostgresRepository.CheckEmailUnique(ctx, logData.Email)
		if err != nil {
			return nil, err
		}
	}
	if !isUsernameExist && !isEmailExist {
		logger.Error("user not exist")
		return nil, errors.NewUserNotExistError("user not exist")
	}

	passwordHash, err := u.userPostgresRepository.GetPasswordHash(ctx, logData.Username, logData.Email)
	if err != nil {
		logger.Error("failed to get password hash")
		return nil, err
	}

	isEqual := CheckPasswordHash(passwordHash, logData.Password)
	if !isEqual {
		logger.Error("wrong password was entered")
		return nil, errors.NewWrongPasswordError("wrong password was entered")
	}

	userID, err := u.userPostgresRepository.GetUserIDByUsername(ctx, logData.Username)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Id:       userID,
		Username: logData.Username,
		Email:    logData.Email,
	}

	return user, nil
}
