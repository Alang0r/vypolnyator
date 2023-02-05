package test

import (
	"testing"

	"github.com/Alang0r/vypolnyator/pkg/log"
	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/sklad/api"
)

func TestSendReq(t *testing.T) {

	l := log.NewLogger()
	l.Init("Tester")
	s := service.NewRequestSender(&l)

	req := api.ReqGroupNew{
		Name:        "First",
		Description: "First group for test",
		UserID:      1,
	}
	rpl := api.ResponseGroupNew{}

	if errReq := s.SendRequest(&req, &rpl); errReq != nil {
		l.Errorf(errReq.Description)
	} else {
		l.Infof("\nSucess! Response: %+v", rpl)
	}

}
