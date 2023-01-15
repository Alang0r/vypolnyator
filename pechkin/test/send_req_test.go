package test

import (
	"log"
	"testing"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/service"
)

type TestReq struct {
	Id   int
	Name string
}

func (r TestReq) Request() string {
	return "/sklad/testRequest"
}

func (r TestReq) Url() string {
	return "http://localhost:3001"
}

type TestRpl struct {
	Exists bool
}

func (r TestReq) Execute() (*TestRpl, *error.Error) {
	rpl := &TestRpl{}
	if r.Id != 1 {
		return nil, error.New().SetCode(error.ErrCodeNotFound)

	}
	rpl.Exists = true
	return rpl, nil

}

func TestSendReq(t *testing.T) {

	s := service.NewRequestSender()

	req := TestReq{}
	req.Id = 1
	req.Name = "Ololow"
	rpl := TestRpl{}

	if errReq := s.SendRequestV2(req, rpl); errReq != nil {
		log.Println(errReq)
	} else {
		log.Println("Sucess!")
	}

}
