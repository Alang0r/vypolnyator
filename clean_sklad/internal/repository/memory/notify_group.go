package memory

import (
	"context"
	"sync"
	"time"

	"github.com/Alang0r/vypolnyator/clean_sklad/internal/entity"
)

type NotifyGroupMemoryStorage struct {
	groups map[uint64]entity.NotifyGroup
	mutex  *sync.Mutex
}

func NewNotifyGroupMemoryStorage() *NotifyGroupMemoryStorage {
	return &NotifyGroupMemoryStorage{
		groups: make(map[uint64]entity.NotifyGroup),
		mutex:  new(sync.Mutex),
	}
}

func (s *NotifyGroupMemoryStorage) CreateNotifyGroup(ctx context.Context, userID uint64, name string, descr string, notifyTime time.Time) error {
	id := uint64(time.Now().UnixNano())
	g := entity.NotifyGroup{
		UserID:      userID,
		Name:        name,
		Description: descr,
		NotifyTime:  notifyTime,
	}
	s.mutex.Lock()
	s.groups[id] = g
	s.mutex.Unlock()
	return nil
}

func (s *NotifyGroupMemoryStorage) GetNotifyGroup(ctx context.Context, id uint64) (*entity.NotifyGroup, error) {
	s.mutex.Lock()
	out := &entity.NotifyGroup{}
	for _, g := range s.groups {
		if g.ID == id {
			out = &g
		}
	}
	s.mutex.Unlock()
	return out, nil
}

func (s *NotifyGroupMemoryStorage) GetNotifyGroups(ctx context.Context) ([]*entity.NotifyGroup, error) {
	out := make([]*entity.NotifyGroup, 0)
	s.mutex.Lock()
	for _, g := range s.groups {
		out = append(out, &g)
	}
	return out, nil
}

func (s *NotifyGroupMemoryStorage) DeleteNotifyGroup(ctx context.Context, id uint64) error {
	s.mutex.Lock()
	for _, g := range s.groups {
		if g.ID == id {
			delete(s.groups, id)
		}
	}
	return nil
}
