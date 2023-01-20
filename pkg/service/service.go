package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"

	err "github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/Alang0r/vypolnyator/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var Handlers map[string]Request
var NewHandlers map[string]reflect.Type
var TypeRegistry = make(map[string]reflect.Type)
var HandlersV2 map[string]RequestV2

func init() {
	Handlers = make(map[string]Request)
	HandlersV2 = make(map[string]RequestV2)
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

func RegisterRequestV2(name string, req RequestV2) {
	HandlersV2[name] = req
}

func RegisterType(typedNil interface{}) {
	t := reflect.TypeOf(typedNil).Elem()
	//TypeRegistry[t.PkgPath()+"."+t.Name()] = t
	TypeRegistry[t.Name()] = t
}

func makeInstance(name string) interface{} {
	return reflect.New(TypeRegistry[name]).Elem().Interface()
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

type UniversalDTO struct {
	Data interface{} `json:"data"`
	// more fields with important meta-data about the message...
}

func SendRequestV1(r Request, url string) (string, []byte) {
	dtoToSend := UniversalDTO{r}
	byteData, _ := json.Marshal(dtoToSend)

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(byteData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	return response.Status, body
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

	// for reqName, req := range Handlers {
	// 	rtGroup.POST(reqName, func(c *gin.Context) {
	// 		err := execRequest(c, reqName, req)
	// 		fmt.Println(err)
	// 		_, _ = req.Execute(c)
	// 		//execRequest(c, reqName, req)
	// 	})
	// }

	rtGroup.GET("")

	log.Printf("Sklad is listening on port: %s", srv.listenAddr)
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

	jsonData, _ := ioutil.ReadAll(c.Request.Body)

	// Create a new struct using reflect
	rType := reflect.TypeOf(Handlers[reqName])
	newStruct := reflect.New(rType).Elem().Interface()
	_ = json.Unmarshal(jsonData, &newStruct)

	// for _, p := range c.Params {
	// 	fmt.Println(p)
	// }
	// tmp1 := makeInstance(reqName)
	// fmt.Println(tmp1)
	// c.Bind(tmp1)
	// fmt.Println(tmp1)

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

func execRequestV2(c *gin.Context, rName string, r RequestV2) err.Error {
	
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(jsonData, &r)

	rpl, err := r.Execute()

	c.JSON(err.GetHttpCode(), gin.H{
		"response": rpl,
	})
	return *err.SetCode(0)
}
