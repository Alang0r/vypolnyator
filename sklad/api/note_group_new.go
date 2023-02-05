package api

import (
	"time"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/sklad/models"
)

func init() {
	service.RegisterHandler("/group/new", (*ReqGroupNew)(nil))
}

type ReqGroupNew struct {
	skladDefault
	Name        string
	Description string
	UserID      uint64
	NotifyTime  time.Time
}

type ResponseGroupNew struct {
	ID uint64
}

func (obj *ReqGroupNew) Request() string {
	return reqPrefix + "/group/new"
}

func (r *ReqGroupNew) Run() (service.Reply, error.Error) {
	l := r.l
	db := r.db

	rpl := ResponseGroupNew{}
	n := models.NoteGroup{
		Name:        r.Name,
		Description: r.Description,
		UserID:      r.UserID,
		NotifyTime:  r.NotifyTime,
	}

	if err := db.Create(&n).Error; err != nil {
		l.Errorf(err.Error())
	}
	rpl.ID = n.ID
	return rpl, *error.New().SetCode(error.ErrCodeNone)
}
