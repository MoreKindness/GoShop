package dal

import (
	"context"
	"gomall/dal/dao"
	"gomall/model"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrRecordNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (repo *UserRepository) Create(ctx context.Context, u model.User) error {
	return repo.dao.Insert(ctx, dao.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	})
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}
	return repo.toModel(u), nil
}

func (repo *UserRepository) toModel(u dao.User) model.User {
	return model.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}
