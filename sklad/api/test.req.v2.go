package api

import (
	"fmt"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
)

func init() {
	service.RegisterHandler("/zTestRequest", &TestReq{})
}

func (r TestReq) Request() string {
	return reqPrefix + "/zTestRequest"
}

type TestReq struct {
	skladDefaultValues
	Id   int
	Name string
}

type TestRpl struct {
	Data string `json:"response"`
}

func (r *TestReq) Execute() (service.Reply, error.Error) {
	l := r.log()
	
	rpl := &TestRpl{}

	rpl.Data = fmt.Sprintf("Privet, %s, tvoy id: %d", r.Name, r.Id)
	l.Info(rpl.Data)

	return rpl, *error.New().SetCode(error.ErrCodeNone)
}
