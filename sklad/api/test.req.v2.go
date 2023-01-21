package api

import (
	"fmt"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
)

func init() {
	service.RegisterRequest("TestRequest",&TestReq{})
}

func (r TestReq) Request() string {
	return "/sklad/TestRequest"
}

func (r TestReq) Url() string {
	return "http://localhost:3001"
}

type TestReq struct {
	Id   int
	Name string
}

type TestRpl struct {
	Data string `json:"response"`
}

func (r *TestReq) Execute() (service.Reply, error.Error) {
rpl := &TestRpl{}

rpl.Data = fmt.Sprintf("Privet, %s, tvoy id: %d",r.Name, r.Id)





	return rpl, *error.New().SetCode(error.ErrCodeNone)
}
