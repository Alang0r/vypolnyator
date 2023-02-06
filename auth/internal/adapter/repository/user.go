package repository

import (
	"context"

	"github.com/Alang0r/vypolnyator/auth/internal/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, id uint64, tgID uint64) error
	GetUser(ctx context.Context, id uint64) (*entity.User, error)
	DeleteUser(ctx context.Context, id uint64) error
}
