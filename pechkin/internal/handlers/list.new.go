package handlers

import (
	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/pkg/telegram"
	"github.com/Alang0r/vypolnyator/sklad/api"
)

func init() {
	telegram.RegisterHandler("/test", (*reqHandlerNewList)(nil))
}

type reqHandlerNewList struct {
	pechkinDefaultValues
}

type rplHanderNewList struct {
	Code    string
	Message string
}

func (h *reqHandlerNewList) Run() (string, error.Error) {
	l := h.Log
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
