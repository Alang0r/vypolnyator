package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Alang0r/vypolnyator/pkg/error"
)

// type Request interface {
// 	Execute(c *gin.Context) (Reply, error.Error)
// }

type Reply interface {
}

type Request interface {
	Request() string
	Url() string
	Execute() (Reply, error.Error)
}

type Response interface {
}

type RequestSender interface {
	SendRequest(Request, Response) *error.Error
}

type Sender struct {
}

func NewRequestSender() Sender {
	s := Sender{}
	return s
}

func (s *Sender) SendRequest(req Request, rpl Response) *error.Error {
	json_data, err := json.Marshal(req)

	if err != nil {
		return error.New().SetCode(error.ErrCodeInternal).SetMessage(err.Error())
	}

	resp, err := http.Post(req.Url()+req.Request(), "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		return error.New().SetCode(error.ErrCodeInternal).SetMessage(err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &rpl)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
