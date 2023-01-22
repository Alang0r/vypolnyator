package api

import(
	"github.com/Alang0r/vypolnyator/pkg/storage"
	"github.com/Alang0r/vypolnyator/pkg/log"
)

func init() {

}

const reqPrefix = "/sklad"

type skladDefaultValues struct {
}

func (v *skladDefaultValues) Url() string {
	return "http://localhost:3001"
}

func (v *skladDefaultValues) storage() storage.Storage {
	s := storage.NewMemoryStorage()
	return s
}

func  (v *skladDefaultValues) log() log.Logger {
	return log.NewLogger()
}

// func  (v *skladDefaultValues) Request() string {
// 	return  "/sklad" + reqName
// }
