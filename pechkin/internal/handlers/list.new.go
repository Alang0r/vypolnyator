package handlers

import (
	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/pkg/telegram"
	"github.com/gin-gonic/gin"
)

func init() {
	telegram.RegisterHandler("/newList", (*HandlerNewList)(nil))
}

type HandlerNewList struct {
}

func (h *HandlerNewList) Execute(c *gin.Context) (service.Reply, error.Error) {

	return nil, *error.New().SetCode(0)
}
