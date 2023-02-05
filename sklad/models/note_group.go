package models

import "time"

type NoteGroup struct {
	ID          uint64
	Name        string
	Description string
	UserID      uint64
	NotifyTime  time.Time
}

func NewNoteGroup(name string, desc string, userID uint64, notifyTime time.Time) {
	
}