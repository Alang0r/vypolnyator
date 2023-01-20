package test

import (
	"log"
	"testing"

	"github.com/Alang0r/vypolnyator/pkg/service"
	"github.com/Alang0r/vypolnyator/sklad/api"
)

// type TestReq struct {
// 	Id   int
// 	Name string
// }

// func (r TestReq) Request() string {
// 	return "/sklad/testRequest"
// }

// func (r TestReq) Url() string {
// 	return "http://localhost:3001"
// }

// type TestRpl struct {
// 	Exists bool
// }

// func (r TestReq) Execute() (*TestRpl, *error.Error) {
// 	rpl := &TestRpl{}
// 	if r.Id != 1 {
// 		return nil, error.New().SetCode(error.ErrCodeNotFound)

// 	}
// 	rpl.Exists = true
// 	return rpl, nil

// }

func TestSendReq(t *testing.T) {

	s := service.NewRequestSender()

	req := api.TestReqV2{}
	req.Id = 1
	req.Name = "Ololow"
	rpl := api.TestRplV2{}

	if errReq := s.SendRequestV2(&req, &rpl); errReq != nil {
		log.Println(errReq)
	} else {
		log.Printf("Sucess! Response: %+v", rpl)
	}

}
