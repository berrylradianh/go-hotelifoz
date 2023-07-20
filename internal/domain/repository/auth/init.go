package auth

import (
	eu "hotelifoz/internal/domain/entity/user"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *eu.User) error
	GetUserByEmail(email string) (*eu.User, error)
}

type authRepo struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) AuthRepository {
	return &authRepo{
		DB,
	}
}
