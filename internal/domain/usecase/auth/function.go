package auth

import (
	eu "hotelifoz/internal/domain/entity/user"
	ra "hotelifoz/internal/domain/repository/auth"
)

type AuthUsecase interface {
	Register(user *eu.User) error
	Login(email, password string) (*eu.User, error)
}

type authUsecase struct {
	authRepository ra.AuthRepository
}

func New(authRepository ra.AuthRepository) *authUsecase {
	return &authUsecase{
		authRepository,
	}
}
