package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/Alang0r/vypolnyator/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var Handlers map[string]Request
var NewHandlers map[string]reflect.Type

func init() {
	Handlers = make(map[string]Request)
	NewHandlers = make(map[string]reflect.Type)
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

}

func RegisterRequest(reqName string, req Request) {
	Handlers[reqName] = req
	t := reflect.TypeOf(req).Elem()
	NewHandlers[reqName] = t
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

func SendRequest(Request) {
/*
	var jsonData = []byte(`{
		"name": "morpheus",
		"job": "leader"
	}`)
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

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

*/
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
		rtGroup.GET(reqName, func(c *gin.Context) {
			execRequest(c, reqName, req)
		})
	}

	rtGroup.GET("")

	srv.router.Run(srv.listenAddr)

}

func (srv *Service) ProcessRequest(req http.Request) error {
	// Get request from

	return nil
}

func (srv *Service) handlerGetPersonById(w http.ResponseWriter, r *http.Request) {

	reqName := r.URL.String()
	log.Printf("Request accepted: %s", reqName)
	resp := http.Response{}
	if _, ok := Handlers[reqName]; ok {
		resp.StatusCode = http.StatusAccepted
		resp.Status = "Request exists!"

	} else {

		resp.StatusCode = http.StatusBadRequest
		resp.Status = fmt.Sprintf("Error: req not in map: %s", reqName)

	}

	json.NewEncoder(w).Encode(resp)

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

func execRequest(c *gin.Context, reqName string, req Request) error {
	// make map with parameters from request
	// fill struct fields from map
	// validate fields
	// execute request
	// return reply

	// Get parameters from request
	params := c.Request.URL.Query()

	tmp := reflect.New(NewHandlers[reqName]).Elem().Interface()

	for param, value := range params {
		err := SetField(tmp, param, value)
		if err != nil {
			return err
		}
	}

	rpl, err := req.Execute(c)

	c.JSON(err.GetHttpCode(), gin.H{
		"response": rpl,
	})
	return nil
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}
