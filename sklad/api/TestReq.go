package api

import (
	"fmt"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/gin-gonic/gin"
)

func init() {
	service.RegisterRequest("testRequest", (*TestReq)(nil))
}

type TestReq struct {
}

type TestRpl struct {
}

func (r *TestReq) Execute(c *gin.Context) (service.Reply, error.Error) {
	fmt.Println("zdarova")
	rpl := TestRpl{}
	return rpl, *error.New().SetCode(0)
}
