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

func (s PostgresStorage) Connect(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	s.db = db
	return nil
}


func (s PostgresStorage) DB() gorm.DB{
	return *s.db
}