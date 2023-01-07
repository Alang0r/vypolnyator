package api

import (
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/gin-gonic/gin"
)

func init() {
	service.RegisterRequest("ReqNewList", (*ReqNewList)(nil))
}

type ReqNewList struct {
}

func (r *ReqNewList) Execute(c *gin.Context) (service.Reply, error.Error) {
	
	return nil, *error.New().SetCode(0)
}
