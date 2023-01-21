package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/gin-gonic/gin"
)

type Request interface {
	Execute(c *gin.Context) (Reply, error.Error)
}

type Reply interface {
}

type RequestV2 interface {
	Request() string
	Url() string
	Execute()  (Reply, error.Error)
}

type Response interface {
}

type RequestSender interface {
	SendRequest(RequestV2, Response) *error.Error
}

type Sender struct {
}

func NewRequestSender() Sender {
	s := Sender{}
	return s
}

func (s *Sender) SendRequest(req RequestV2, rpl Response) *error.Error {

	resp, err := http.Get(req.Url() + req.Request())

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	return nil
}

func (s *Sender) SendRequestV2(req RequestV2, rpl Response) *error.Error {
	json_data, err := json.Marshal(req)

	if err != nil {
		return error.New().SetCode(error.ErrCodeInternal).SetMessage(err.Error())
	}

	resp, err := http.Post(req.Url()+req.Request(), "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		return error.New().SetCode(error.ErrCodeInternal).SetMessage(err.Error())
	}

	// str := make(map[string]interface{})
	// 	body, err := ioutil.ReadAll(resp.Body)
	// json.Unmarshal(body, &str)
	// fmt.Println(str)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	err = json.Unmarshal(body, &rpl)
	if err != nil {
        fmt.Println(err)
    }


	//json.NewDecoder(resp.Body).Decode(&rpl)
	return nil
}
