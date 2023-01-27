package api

import (
	"fmt"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
)

func init() {
	//reqName =
	service.RegisterHandler("/list/new", &RequestListNew{})
}

func (h *RequestListNew) Request() string {
	return reqPrefix + "/list/new"
}

type RequestListNew struct {
	skladDefaultValues
	Name string
}

type ResponseListNew struct {
	ID    uint64
	Hello string
}

func (r *RequestListNew) Run() (service.Reply, error.Error) {
	rpl := ResponseListNew{}
	rpl.ID = 666
	rpl.Hello = fmt.Sprintf("Privet iz %s tebya zovut %s", "list.new", r.Name)

	return rpl, *error.New().SetCode(error.ErrCodeNone)
}
