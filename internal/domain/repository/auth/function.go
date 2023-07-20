package auth

import (
	"errors"

	eu "hotelifoz/internal/domain/entity/user"
)

func (ar *authRepo) CreateUser(user *eu.User) error {
	existingUser := &eu.User{}
	err := ar.DB.Where("email = ?", user.Email).First(existingUser).Error
	if err != nil {
		err = ar.DB.Create(&user).Error
		if err != nil {
			return err
		}
	} else {
		//lint:ignore ST1005 Reason for ignoring this linter
		return errors.New("Email already exists")
	}

	return nil
}

func (ar *authRepo) GetUserByEmail(email string) (*eu.User, error) {
	user := &eu.User{}
	err := ar.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, errors.New("Record Not Found")
	}

	return user, nil
}
