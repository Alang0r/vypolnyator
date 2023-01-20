package api

import (
	"fmt"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/gin-gonic/gin"
)

func init() {
	service.RegisterRequest("TestRequest", (*TestReq)(nil))
	service.RegisterType( (*TestReq)(nil))
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
	Data string
}

func (r *TestReq) Execute(c *gin.Context) (service.Reply, error.Error) {
	fmt.Println("zdarova")
	rpl := TestRpl{}
	rpl.Data = fmt.Sprintf("Helo, ")
	return rpl, *error.New().SetCode(0)
}
