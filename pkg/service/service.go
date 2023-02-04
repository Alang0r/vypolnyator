package service

import (
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/Alang0r/vypolnyator/pkg/log"
	"github.com/Alang0r/vypolnyator/pkg/middleware"
	"github.com/Alang0r/vypolnyator/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var Handlers map[string]Handler

func init() {
	Handlers = make(map[string]Handler)
}

func RegisterHandler(name string, h Handler) {
	Handlers[name] = h
}

type Service struct {
	name       string
	listenAddr string
	store      storage.Storage
	router     *gin.Engine
	Log        log.Logger
	Params     map[string]string
}

func NewService(serviceName string, listenAddr string, storage storage.Storage) *Service {
	l := log.NewLogger()
	l.Init(serviceName)
	return &Service{
		name:       serviceName,
		store:      storage,
		listenAddr: listenAddr,
		Params:     make(map[string]string),
		Log:        l,
	}
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

/*
func execRequest(c *gin.Context, rName string, r Handler) err.Error {

	jsonData, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(jsonData, &r)

	rpl, error := r.Run()
	c.JSON(error.GetHttpCode(), rpl)
	return *err.New().SetCode(0)
}

func execRequestV2(c *gin.Context) err.Error {

	return *err.New().SetCode(err.ErrCodeNone)
}
*/

func (srv *Service) Listen() {
	// set gin
	gin.SetMode(gin.ReleaseMode)
	srv.router = gin.New()

	// set middleware
	srv.router.Use(middleware.NewMiddleware(&srv.Log))
	srv.router.Use(gin.Recovery())

	// add handlers
	for reqName, req := range Handlers {
		hndlr := reflect.New(reflect.TypeOf(req).Elem()).Interface().(Handler)
		//h := req
		hndlr.SetEnv(&srv.Log, srv.store.DB())
		srv.router.POST(srv.name+reqName, func(c *gin.Context) {
			c.BindJSON(&hndlr)
			rpl, err := hndlr.Run()
			c.JSON(err.GetHttpCode(), rpl)
		})
	}

	srv.Log.Infof("listening on port %s", srv.listenAddr)
	srv.router.Run(srv.listenAddr)
}

func (s *Service) GetEnvVariable(name string) string {
	// Get env variables
	err := godotenv.Load()
	if err != nil {
		s.Log.Fatalf("Error loading .env file")
	}
	return os.Getenv(name)
}

func (s *Service) SetDB() error {
	host := s.GetEnvVariable(storage.ParamDBHOST)
	port := s.GetEnvVariable(storage.ParamDBPORT)
	user := s.GetEnvVariable(storage.ParamDBUser)
	pass := s.GetEnvVariable(storage.ParamDBPass)
	dbName := s.GetEnvVariable(storage.ParamDBName)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow", host, user, pass, dbName, port)

	if err := s.store.Connect(dsn); err != nil {
		s.Log.Errorf("error connect to DB %s: %s", dbName, err.Error())
		return err
	}

	// if err := storage.Connect(dsn);err != nil {
	// 	s.Log.Errorf("error connect to DB %s: %s", dbName, err.Error())
	// 	return err
	// }
	s.Log.Infof("successfully conected to DB %s", dbName)
	return nil
}
