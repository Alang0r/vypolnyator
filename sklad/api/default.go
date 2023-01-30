package api

import(
	"github.com/Alang0r/vypolnyator/pkg/storage"
	"github.com/Alang0r/vypolnyator/pkg/log"
)

func init() {

}

const reqPrefix = "/sklad"

type skladDefaultValues struct {
	l *log.Logger
}

func (v *skladDefaultValues) Url() string {
	return "http://localhost:3001"
}

func (v *skladDefaultValues) SetLog(l *log.Logger) {
	v.l = l
}

func (v *skladDefaultValues) storage() storage.Storage {
	s := storage.NewMemoryStorage()
	return s
}

func (v *skladDefaultValues) Log() *log.Logger {
	return v.l
}

// func  (v *skladDefaultValues) log() log.Logger {
// 	return *v.l
// }

// func  (v *skladDefaultValues) Request() string {
// 	return  "/sklad" + reqName
// }
