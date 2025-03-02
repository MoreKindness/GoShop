package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gomall/dal"
	"gomall/dal/dao"
	"gomall/model"
)

var (
	ErrDuplicateEmail        = dao.ErrDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("用户不存在或者密码不对")
)

type UserService struct {
	repo *dal.UserRepository
}

func NewUserService(repo *dal.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Signup(ctx context.Context, u model.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, email string, password string) (model.User, error) {
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == dal.ErrUserNotFound {
		return model.User{}, err
	}
	if err != nil {
		return model.User{}, err
	}

	//	检查密码是否中正确
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return model.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}
