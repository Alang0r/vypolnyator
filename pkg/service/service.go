package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	err "github.com/Alang0r/vypolnyator/pkg/error"
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
	return &Service{
		name:       serviceName,
		store:      storage,
		listenAddr: listenAddr,
		Params:     make(map[string]string),
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

func (srv *Service) Listen() {
	// Init log
	srv.Log.Init(srv.name)
	// Setting gin
	gin.SetMode(gin.ReleaseMode)
	srv.router = gin.New()

	srv.router.Use(middleware.NewMiddleware(&srv.Log))
	for reqName, req := range Handlers {
		h := req
		srv.router.POST(srv.name+reqName, func(c *gin.Context) {
			c.BindJSON(&h)
			rpl, err := h.Run()
			c.JSON(err.GetHttpCode(), rpl)

		})
	}

	srv.Log.Infof("listening on port: %s", srv.listenAddr)
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
