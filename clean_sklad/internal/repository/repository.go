package repository

import (
	"context"
	"time"

	"github.com/Alang0r/vypolnyator/clean_sklad/internal/entity"
)

type NotifyGroupRepository interface {
	CreateNotifyGroup(ctx context.Context, userID uint64, name string, descr string, notifyTime time.Time) error
	GetNotifyGroup(ctx context.Context, id uint64) (*entity.NotifyGroup, error)
	GetNotifyGroups(ctx context.Context) ([]*entity.NotifyGroup, error)
	DeleteNotifyGroup(ctx context.Context, id uint64) error
}

