package api

import(
	"github.com/Alang0r/vypolnyator/pkg/storage"
	"github.com/Alang0r/vypolnyator/pkg/log"
)

func init() {

}

const reqPrefix = "/sklad"

type skladDefault struct {
	l *log.Logger
	reqID string
}

func (v *skladDefault) Url() string {
	return "http://localhost:3001"
}

func (v *skladDefault) SetLog(l *log.Logger) {
	v.l = l
}

func (v *skladDefault) SetReqID(id string) {
	v.reqID = id
}

func (v *skladDefault) GetReqID() string {
	return v.reqID
}

func (v *skladDefault) storage() storage.Storage {
	s := storage.NewMemoryStorage()
	return s
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
