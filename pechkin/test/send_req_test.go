package test

import (
	"log"
	"testing"

	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/sklad/api"
)

func TestSendReq(t *testing.T) {

	s := service.NewRequestSender()

	req := api.TestReq{}
	req.Id = 5
	req.Name = "Vasya"
	rpl := api.TestRpl{}

	if errReq := s.SendRequest(&req, &rpl); errReq != nil {
		log.Println(errReq)
	} else {
		log.Printf("Sucess! Response: %+v", rpl)
	}

}
