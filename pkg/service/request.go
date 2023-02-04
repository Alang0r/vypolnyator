package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reply interface {
}

type Handler interface {
	Request() string
	Url() string
	Run() (Reply, error.Error)
	SetEnv(*log.Logger, gorm.DB)
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
		s.l.Infof(err.Error())
	}

	return nil
}

func (s *Sender) SendRequestV1(req Handler, rpl Response) *error.Error {
	// marshal req to json
	jsonData, err := json.Marshal(req)
	if err != nil {
		return error.New().SetCode(error.ErrCodeInternal).SetMessage(err.Error())
	}

	reqName := req.Request()
	reqID := req.GetReqID()
	if reqID == "" {
		reqID = generateReqID()
		req.SetReqID(reqID)
	}
	url := req.Url() + reqName

	// create request
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return error.New().SetCode(error.ErrCodeInternal).SetMessage(err.Error())
	}

	// set headers
	//request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("ReqID", reqID)
	request.Header.Set("Request", reqName)

	// send request
	s.l.Infof("Sending %s, data: %s,  id <%s> to <%s>",string(jsonData), reqName, reqID, req.Url())
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		s.l.Errorf("error sending request", err.Error())
	}

	defer resp.Body.Close()

	// resp, err := http.Post(req.Url()+req.Request(), "application/json",
	// 	bytes.NewBuffer(jsonData))

	if err != nil {
		return error.New().SetCode(error.ErrCodeInternal).SetMessage(err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.l.Errorf("error sending request", err.Error())
	}

	
	s.l.Infof("Get response: %+v", string(body))
	err = json.Unmarshal(body, &rpl)
	if err != nil {
		s.l.Errorf("error sending request, %s", err.Error())
	}

	return nil
}

// generateReqID - generates uniq requestID
func generateReqID() string {
	return uuid.New().String()
}
