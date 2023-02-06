package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/Alang0r/vypolnyator/auth/internal/entity"
)

type UserMemoryStorage struct {
	users map[uint64]entity.User
	mutex *sync.Mutex
}

func NewUserMemorytorage() *UserMemoryStorage {
	return &UserMemoryStorage{
		users: make(map[uint64]entity.User),
		mutex: new(sync.Mutex),
	}
}

func (s *UserMemoryStorage) CreateUser(ctx context.Context, id uint64, tgID uint64) error {
	u := entity.User{
		ID: id,
		TelegramID: tgID,
	}
	s.mutex.Lock()
	s.users[id] = u
	s.mutex.Unlock()

	return nil
}

func (s *UserMemoryStorage) GetUser(ctx context.Context, id uint64) (*entity.User, error) {
	s.mutex.Lock()
	out := &entity.User{}
	for _, u := range s.users {
		if u.ID == id {
			out = &u
		}
	}
	s.mutex.Unlock()
	return out, nil
}

func (s *UserMemoryStorage) DeleteUser(ctx context.Context, id uint64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	u, ex := s.users[id]
	if ex && u.ID == id {
		delete(s.users, id)
		return nil
	}

	return fmt.Errorf("user not found")
}
