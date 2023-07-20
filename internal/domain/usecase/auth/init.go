package auth

import (
	eu "hotelifoz/internal/domain/entity/user"
	hlp "hotelifoz/internal/domain/helper"
)

func (au *authUsecase) Register(user *eu.User) error {
	hashedPassword, err := hlp.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = au.authRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (au *authUsecase) Login(email, password string) (*eu.User, error) {
	user, err := au.authRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = hlp.VerifyPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
