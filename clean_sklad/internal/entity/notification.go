package entity

import "time"

type Notification struct {
	ID            uint64
	NotifyGroupID uint64
	Name          string
	Description   string
	NotifyTime    time.Time
}
