package storage

import "gorm.io/gorm"

type Storage interface {
	Connect(string) error
	DB() gorm.DB
}
