package test

import (
	"log"
	"testing"

	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/sklad/api"
)

func TestSendReq(t *testing.T) {

	s := service.NewRequestSender()

	req := api.TestReqV2{}
	req.Id = 5
	req.Name = "Sasha"
	rpl := api.TestRplV2{}

	if errReq := s.SendRequestV2(&req, &rpl); errReq != nil {
		log.Println(errReq)
	} else {
		log.Printf("Sucess! Response: %+v", rpl)
	}

}
