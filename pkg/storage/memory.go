package storage

import "gorm.io/gorm"

type MemoryStorage struct {
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(uint64) {
}

func (s *MemoryStorage) Connect(string) error {
	return nil
}
func (s *MemoryStorage) DB() gorm.DB {
	return gorm.DB{}
}
