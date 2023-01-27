package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Alang0r/vypolnyator/pkg/error"
)

type Reply interface {
}

type Handler interface {
	Request() string
	Url() string
	Run() (Reply, error.Error)
}

type Response interface {
}

type RequestSender interface {
	SendRequest(Handler, Response) *error.Error
}

type Sender struct {
}

func NewRequestSender() Sender {
	s := Sender{}
	return s
}

func (s *Sender) SendRequest(req Handler, rpl Response) *error.Error {
	json_data, err := json.Marshal(req)

	if err != nil {
		return error.New().SetCode(error.ErrCodeInternal).SetMessage(err.Error())
	}

	fmt.Printf("Sending <%s> request to <%s>", req.Request(), req.Url())
	resp, err := http.Post(req.Url()+req.Request(), "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		return error.New().SetCode(error.ErrCodeInternal).SetMessage(err.Error())
	}

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &rpl)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
