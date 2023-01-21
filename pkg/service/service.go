package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	err "github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/storage"
	"github.com/gin-gonic/gin"
)

var Handlers map[string]Request

func init() {
	Handlers = make(map[string]Request)
}
func RegisterRequest(name string, req Request) {
	Handlers[name] = req
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
	for reqName, req := range Handlers {
		rtGroup.POST(reqName, func(c *gin.Context) {
			execRequest(c, reqName, req)
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

func execRequest(c *gin.Context, rName string, r Request) err.Error {

	jsonData, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(jsonData, &r)

	rpl, err := r.Execute()
	c.JSON(err.GetHttpCode(), rpl)
	return *err.SetCode(0)
}
