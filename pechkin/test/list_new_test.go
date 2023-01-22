package test

import (
	"fmt"
	"testing"

	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/sklad/api"
)

func TestListNew(t *testing.T) {

	s := service.NewRequestSender()

	req := api.RequestListNew{}
	req.Name = "Dima"
	rpl := api.ResponseListNew{}

	if errReq := s.SendRequest(&req, &rpl); errReq != nil {
		fmt.Println(errReq)
	} else {
		fmt.Printf("\nSucess! Response: %+v", rpl)
	}

}
