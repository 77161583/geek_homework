package service

import (
	"context"
	"home2/internal/domain"
	"home2/internal/repository"
)

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
	//2.存起来
	return svc.repo.Created(ctx, u)
}
