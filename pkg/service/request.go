package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/log"
	"github.com/google/uuid"
)

type Reply interface {
}

type Handler interface {
	Request() string
	Url() string
	Run() (Reply, error.Error)
	SetLog(*log.Logger)
	Log() *log.Logger
	SetReqID(string)
	GetReqID() string
}

type Response interface {
}

type RequestSender interface {
	SendRequest(Handler, Response) *error.Error
}

type Sender struct {
	l *log.Logger
}

func NewRequestSender(l *log.Logger) Sender {
	s := Sender{
		l: l,
	}

	return s
}

func (s *Sender) SendRequest(req Handler, rpl Response) *error.Error {
	json_data, err := json.Marshal(req)

	if err != nil {
		return error.New().SetCode(error.ErrCodeInternal).SetMessage(err.Error())
	}

	if req.GetReqID() == "" {
		req.SetReqID(generateReqID())
	}

	s.l.Infof("Sending <%s> request with id <%s> to <%s>", req.Request(), req.GetReqID(), req.Url())
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

func generateReqID() string {
	return uuid.New().String()
}
