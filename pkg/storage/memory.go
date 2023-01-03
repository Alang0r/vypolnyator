package storage

type MemoryStorage struct {
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(uint64) {
}
