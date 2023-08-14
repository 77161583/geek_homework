package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"home2/internal/domain"
	"home2/internal/repository"
)

var ErrUseDuplicateEmail = repository.ErrUseDuplicateEmail
var ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码不对")

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

/**
 * SignUp
 * ctx context.Context 为了保持链路和可观测性，所以加了这个
 * 不知道方法返回什么就返回一个error
 */
func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	//1.加密问题
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	//2.存起来
	return svc.repo.Created(ctx, u)
}

func (svc *UserService) Login(ctx context.Context, email, password string) error {
	//先找用户
	u, err := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFund {
		return ErrInvalidUserOrPassword
	}
	if err != nil {
		return err
	}
	//比较密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		//DEBUG
		return ErrInvalidUserOrPassword
	}
	return err
}
