package handlers

import (
	"github.com/Alang0r/vypolnyator/pkg/log"
)


type PechkinDefaultValues struct {
	log *log.Logger
}

func init() {

}

func (v *PechkinDefaultValues) SetLog(l *log.Logger) {
	v.log = l
}


// func (v *PechkinDefaultValues) Log() *log.Logger {
// 	return v.Logger
// }
