package service_V1

import (
	"github.com/Alang0r/vypolnyator/pkg/log"
	"github.com/Alang0r/vypolnyator/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	InitRoutes() *gin.Engine
	SetRoutes(g *gin.RouterGroup)
}

type Service struct {
	name       string
	listenAddr string
	router     *gin.Engine
	Log        log.Logger
}

func NewService(serviceName string, listenAddr string) *Service {
	l := log.NewLogger()
	l.Init(serviceName)
	// set gin
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// set middleware
	router.Use(middleware.NewMiddleware(&l))
	router.Use(gin.Recovery())

	return &Service{
		name:       serviceName,
		listenAddr: listenAddr,
		Log:        l,
		router:     router,
	}
}

func (srv *Service) Listen() {

	srv.Log.Infof("listening on port %s", srv.listenAddr)
	srv.router.Run(srv.listenAddr)
}

func (s *Service) AddHAndlers(h Handler) {
	srv := s.router.Group("/" + s.name)

	h.SetRoutes(srv)

}
