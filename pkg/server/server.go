package server

import (
	"github.com/1ch0/gin-template/pkg/server/config"
	"github.com/1ch0/gin-template/pkg/server/interfaces/api"
	"github.com/1ch0/gin-template/pkg/server/utils/container"
	"github.com/1ch0/gin-template/pkg/server/utils/log"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run() error
}

type restServer struct {
	webContainer  *gin.Engine
	beanContainer *container.Container
	cfg           config.Config
}

func New(cfg config.Config) Server {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(gin.LoggerWithWriter(log.Logger.Writer()))
	return &restServer{
		webContainer: app,
		cfg:          cfg,
	}
}

func (r *restServer) buildIoCContainer() (err error) {
	r.beanContainer = container.NewContainer()

	return nil
}

func (r *restServer) RegisterAPIRoute() {
	api.InitAPIBean(r.webContainer)
	for _, handler := range api.GetRegisteredAPI() {
		handler.Register()
	}
}

func (r *restServer) Run() error {
	if err := r.buildIoCContainer(); err != nil {
		return err
	}
	r.RegisterAPIRoute()
	log.Logger.Infof("Server is running on %s", r.cfg.Server.BindAddr)
	return r.webContainer.Run(r.cfg.Server.BindAddr)
}
