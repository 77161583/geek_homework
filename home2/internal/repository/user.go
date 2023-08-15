package repository

import (
	"context"
	"home2/internal/domain"
	"home2/internal/repository/dao"
)

var ErrUseDuplicateEmail = dao.ErrUseDuplicateEmail
var ErrUserNotFund = dao.ErrUserNotFund

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{Id: u.Id, Email: u.Email, Password: u.Password}, nil
}

func (r *UserRepository) Created(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	return r.dap.
}
