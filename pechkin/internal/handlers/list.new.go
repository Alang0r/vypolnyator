package handlers

import (
	"fmt"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/pkg/telegram"
	"github.com/Alang0r/vypolnyator/sklad/api"
)

func init() {
	telegram.RegisterHandler("/test", (*reqHandlerNewList)(nil))
}

type reqHandlerNewList struct {
}

type rplHanderNewList struct {
	Code    string
	Message string
}

func (h *reqHandlerNewList) Run() (string, error.Error) {

	//out := rplHanderNewList{}
	s := service.NewRequestSender()

	req := api.RequestListNew{}
	req.Name = "Dima"
	rpl := api.ResponseListNew{}

	if errReq := s.SendRequest(&req, &rpl); errReq != nil {
		fmt.Println(errReq)
	} else {
		fmt.Printf("\nSucess! Response: %+v", rpl)
	}

	return rpl.Hello, *error.New().SetCode(0)
}
