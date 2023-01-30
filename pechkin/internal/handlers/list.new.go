package handlers

import (
	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/pkg/telegram"
	"github.com/Alang0r/vypolnyator/sklad/api"
)

func init() {
	telegram.RegisterHandler("/test", (*ReqHandlerNewList)(nil))
}

type ReqHandlerNewList struct {
	PechkinDefaultValues
	Name string
}

type RplHanderNewList struct {
	Code    string
	Message string
}

func (h *ReqHandlerNewList) Run() (string, error.Error) {
	l := h.log
	s := service.NewRequestSender(l)

	req := api.RequestListNew{}
	req.Name = "Dima"
	rpl := api.ResponseListNew{}

	if errReq := s.SendRequest(&req, &rpl); errReq != nil {
		l.Info(errReq.Description)
	} else {
		l.Infof("Sucess! Response: %+v", rpl)
	}

	return rpl.Hello, *error.New().SetCode(0)
}
