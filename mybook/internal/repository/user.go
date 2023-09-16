package repository

import (
	"context"
	"mybook/internal/domain"
	"mybook/internal/repository/dao"
	"time"
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

func (r *UserRepository) Edit(ctx context.Context, u domain.User) error {
	return r.dao.Edit(ctx, dao.User{
		Id:           u.Id,
		NickName:     u.NickName,
		Birthday:     u.Birthday,
		Introduction: u.Introduction,
	})
}

func (r *UserRepository) FindById(ctx context.Context, userId int64) (domain.User, error) {
	u, err := r.dao.FindById(ctx, userId)
	if err != nil {
		return domain.User{}, err
	}
	t1 := u.CreateTime / 1000
	t2 := u.UpdateTime / 1000
	return domain.User{
		Id:           u.Id,
		Email:        u.Email,
		Password:     u.Password,
		NickName:     u.NickName,
		Birthday:     u.Birthday,
		Introduction: u.Introduction,
		CreateTime:   time.Unix(t1, 0).String(),
		UpdateTime:   time.Unix(t2, 0).String(),
	}, nil
}
