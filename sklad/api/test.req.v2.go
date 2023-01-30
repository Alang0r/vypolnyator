package api

import (
	"fmt"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
)

func init() {
	request := "TestRequest"
	service.RegisterHandler(request, &TestReq{})
}

func (r TestReq) Request() string {
	return reqPrefix + "TestRequest"
}

type TestReq struct {
	skladDefault
	Id   int
	Name string
}

type TestRpl struct {
	Data string `json:"response"`
}

func (r *TestReq) Run() (service.Reply, error.Error) {
	l := r.l

	rpl := &TestRpl{}

	rpl.Data = fmt.Sprintf("Privet, %s, tvoy id: %d", r.Name, r.Id)
	l.Info(rpl.Data)

	return rpl, *error.New().SetCode(error.ErrCodeNone)
}
