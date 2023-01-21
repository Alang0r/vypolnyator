package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	err "github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/storage"
	"github.com/gin-gonic/gin"
)

var HandlersV2 map[string]RequestV2

func init() {
	HandlersV2 = make(map[string]RequestV2)
}
func RegisterRequestV2(name string, req RequestV2) {
	HandlersV2[name] = req
}

type Service struct {
	name       string
	listenAddr string
	store      storage.Storage
	router     *gin.Engine
	Params     map[string]string
}

func NewService(serviceName string, listenAddr string, storage storage.Storage) *Service {
	return &Service{
		name:       serviceName,
		store:      storage,
		listenAddr: listenAddr,
		Params:     make(map[string]string),
	}
}

// SendRequest - просто берем строку и шлем на юрл
func SendRequestV2(reqStr string, url string) string {

	var jsonData = []byte(`{
		"name": "morpheus",
		"job": "leader"
	}`)

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if error != nil {
		fmt.Println(error)
		return ""
	}
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)

	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return string(body)

}

func (srv *Service) Start() {
	srv.router = gin.Default()
	/*
		Тут нужно
		1. брать полученный запрос из мапы
		2. определять тип с помощью reflect
		3. создавать экземпляр структуры
		4. передавать execute экземпляра в группу

		//rtGroup := srv.router.Group("/person")
		// for _, req := range Handlers {
		// 	rtGroup.Handlers = append(rtGroup.Handlers, req)
		// }
		// rtGroup.Handlers = append(rtGroup.Handlers, )

	*/

	rtGroup := srv.router.Group(fmt.Sprintf("/%s", srv.name))
	for reqName, req := range HandlersV2 {
		rtGroup.POST(reqName, func(c *gin.Context) {
			execRequestV2(c, reqName, req)
		})
	}

	log.Printf("Sklad is listening on port: %s", srv.listenAddr)
	srv.router.Run(srv.listenAddr)

}

func (srv *Service) ProcessRequest(req http.Request) error {
	// Get request from

	return nil
}

func (srv *Service) GetParameters(paramName ...string) error {
	for _, pName := range paramName {
		p, err := getEnv(pName)
		if err != nil {
			return err
		}
		srv.Params[pName] = p
	}
	return nil
}

func getEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("parameter not found %s", key)
	}

	return value, nil
}

func execRequestV2(c *gin.Context, rName string, r RequestV2) err.Error {

	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(jsonData, &r)

	rpl, err := r.Execute()
	c.JSON(err.GetHttpCode(), rpl)
	return *err.SetCode(0)
}
