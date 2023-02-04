package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresStorage struct {
	db *gorm.DB
}

func NewPGStorage() PostgresStorage {
	return PostgresStorage{}
}

func (s *PostgresStorage) Connect(dsn string) error {
	var err error
	s.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}


func (s *PostgresStorage) DB() gorm.DB{
	return *s.db
}