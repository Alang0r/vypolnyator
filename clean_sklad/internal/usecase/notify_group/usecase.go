package notify_group

import (
	"context"
	"time"

	"github.com/Alang0r/vypolnyator/clean_sklad/internal/entity"
	"github.com/Alang0r/vypolnyator/clean_sklad/internal/repository"
)


type Usecase interface {
	CreateNotifyGroup(ctx context.Context, userID uint64, name string, descr string, notifyTime time.Time) error
	GetNotifyGroup(ctx context.Context, id uint64) (*entity.NotifyGroup, error)
	GetNotifyGroups(ctx context.Context) ([]*entity.NotifyGroup, error)
	DeleteNotifyGroup(ctx context.Context, id uint64) error
}

type NotifyGroupUseCase struct {
	repository repository.NotifyGroupRepository
}

func NewNotifyGroupUseCase(rep repository.NotifyGroupRepository) *NotifyGroupUseCase {
	return &NotifyGroupUseCase{
		repository: rep,
	}
}

func (u NotifyGroupUseCase) CreateNotifyGroup(ctx context.Context, userID uint64, name string, descr string, notifyTime time.Time) error {
	return u.repository.CreateNotifyGroup(ctx, userID, name, descr, notifyTime)
}

func (u NotifyGroupUseCase) GetNotifyGroup(ctx context.Context, id uint64) (*entity.NotifyGroup, error) {
	return u.repository.GetNotifyGroup(ctx, id)
}

func (u NotifyGroupUseCase) GetNotifyGroups(ctx context.Context) ([]*entity.NotifyGroup, error) {
	return u.repository.GetNotifyGroups(ctx)
}

func (u NotifyGroupUseCase) DeleteNotifyGroup(ctx context.Context, id uint64) error {
	return u.repository.DeleteNotifyGroup(ctx, id)
}
