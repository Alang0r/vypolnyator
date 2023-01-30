package handlers

import (
	"github.com/Alang0r/vypolnyator/pkg/log"
)

type pechkinDefaultValues struct {
	Log *log.Logger
}

func init() {

}

func (v *pechkinDefaultValues) SetLog(l *log.Logger) {
}

// func (v *pechkinDefaultValues) Log() *log.Logger {
// 	return Tlog
// }
