package usecase

import (
	"fmt"
	"strings"

	"github.com/BradleyZhang/ecommerce/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo      domain.UserRepository
	validater domain.Validater
}

func NewUserUsecase(repo domain.UserRepository, v domain.Validater) *UserUsecase {
	return &UserUsecase{
		repo:      repo,
		validater: v,
	}
}

// TODO 包装各种错误，不要把 比如说repo的数据库错误流到delivery
func (u *UserUsecase) Registration(username string, password string) error {
	if password != "" { // 空密码数据库存空
		var err error
		password, err = password2Hash(password)
		if err != nil {
			return err
		}
	}
	username = strings.TrimSpace(username)
	user := domain.User{
		UserName: username,
		Password: password,
		Role:     domain.UserRole,
	}
	if err := u.validater.Struct(&user); err != nil {
		return err
	}
	exist, err := u.repo.Exist(username)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("user exist")
	}

	_, err = u.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func password2Hash(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// func (s *UserService)
