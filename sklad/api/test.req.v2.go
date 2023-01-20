package api

import (
	"fmt"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
)

func init() {
	service.RegisterRequestV2("TestRequestV2",&TestReqV2{})
}

func (r TestReqV2) Request() string {
	return "/sklad/TestRequestV2"
}

func (r TestReqV2) Url() string {
	return "http://localhost:3001"
}

type TestReqV2 struct {
	Id   int
	Name string
}

type TestRplV2 struct {
	Data string `json:"response"`
}

func (r *TestReqV2) Execute() (service.Reply, error.Error) {
rpl := &TestRplV2{}

rpl.Data = fmt.Sprintf("Privet, %s, tvoy id: %d",r.Name, r.Id)





	return rpl, *error.New().SetCode(error.ErrCodeNone)
}
