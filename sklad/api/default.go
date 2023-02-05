package api

import (
	"github.com/Alang0r/vypolnyator/pkg/log"
	"gorm.io/gorm"
)

func init() {

}

const reqPrefix = "/sklad"

type skladDefault struct {
	l     *log.Logger
	db    gorm.DB
	reqID string
}

func (v *skladDefault) Url() string {
	return "http://localhost:3002"
}

func (v *skladDefault) SetEnv(l *log.Logger, db gorm.DB) {
	v.l = l
	v.db = db
}

func (v *skladDefault) SetReqID(id string) {
	v.reqID = id
}

func (v *skladDefault) GetReqID() string {
	return v.reqID
}

func (v *skladDefault) storage() gorm.DB {
	return v.db
}

func (v *skladDefault) Log() *log.Logger {
	return v.l
}

// func  (v *skladDefault) log() log.Logger {
// 	return *v.l
// }

// func  (v *skladDefault) Request() string {
// 	return  "/sklad" + reqName
// }
