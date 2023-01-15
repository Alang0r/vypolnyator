package api

import (
	"fmt"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/gin-gonic/gin"
)

func init() {
	service.RegisterRequest("TestRequest", (*TestRequest)(nil))
	service.RegisterType( (*TestRequest)(nil))
}

func (r TestRequest) Request() string {
	return "/sklad/TestRequest"
}

func (r TestRequest) Url() string {
	return "http://localhost:3001"
}


type TestRequest struct {
	Id   int
	Name string
}

type TestRpl struct {
	Exists bool
}

func (r *TestRequest) Execute(c *gin.Context) (service.Reply, error.Error) {
	fmt.Println("zdarova")
	rpl := TestRpl{}
	return rpl, *error.New().SetCode(0)
}
