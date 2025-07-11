package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"

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

func HashPassword(salt []byte, password string) string {
	hashedPass := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	combined := append(salt, hashedPass...)
	return base64.StdEncoding.EncodeToString(combined)
}

func (u *userUsecase) SignupUser(ctx context.Context, regData *model.RegisterData) (*model.User, error) {
	isUniqueUsername, err := u.userPostgresRepository.CheckUsernameUnique(ctx, regData.Username)
	if err != nil {
		return nil, err
	}
	if !isUniqueUsername {
		return nil, errors.NewNotUniqueError("username %s not unique", regData.Username)
	}
	isUniqueEmail, err := u.userPostgresRepository.CheckEmailUnique(ctx, regData.Email)
	if err != nil {
		return nil, err
	}
	if !isUniqueEmail {
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
