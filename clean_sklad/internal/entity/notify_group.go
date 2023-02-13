package entity

import "time"

type NotifyGroup struct {
	ID          uint64
	UserID      uint64
	Name        string
	Description string
	NotifyTime  time.Time
}
