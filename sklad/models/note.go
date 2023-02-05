package models

import "time"

type Note struct {
	ID          uint64
	Name        string
	Description string
	UserID      uint64
	NotifyTime  time.Time
}

func NewNote(name string, desc string, userID uint64, notifyTime time.Time) {
	
}