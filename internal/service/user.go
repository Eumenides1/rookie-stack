package service

import (
	"context"
	"github.com/Eumenides/rookie-stack/internal/domain"
	"github.com/Eumenides/rookie-stack/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {

	return svc.repo.Create(ctx, u)
}
