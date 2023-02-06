package usecase

import (
	"context"

	"github.com/Alang0r/vypolnyator/auth/internal/adapter/repository"
	"github.com/Alang0r/vypolnyator/auth/internal/entity"
)

type UserUseCase struct {
	userRepo repository.Repository
}

func NewUserUseCase(userRepo repository.Repository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (u UserUseCase) CreateUser(ctx context.Context, id uint64, tgID uint64) error {
	return u.userRepo.CreateUser(ctx, id, tgID)
}

func (u UserUseCase) GetUser(ctx context.Context, id uint64) (*entity.User, error) {
	return u.userRepo.GetUser(ctx, id)
}
func (u UserUseCase) DeleteUser(ctx context.Context, id uint64) error {
	return u.userRepo.DeleteUser(ctx, id)
}
