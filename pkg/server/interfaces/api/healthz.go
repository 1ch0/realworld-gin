package api

import (
	"net/http"

	"github.com/1ch0/gin-template/pkg/server/utils/log"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type healthzApi struct {
	*gin.Engine
}

func NewHealthzApi(app *gin.Engine) Interface {
	return &healthzApi{app}
}

func (h *healthzApi) Register() {
	api := h.Group("/healthz")
	api.GET("", h.healthz)
}

func (h *healthzApi) healthz(c *gin.Context) {
	log.Logger.Log(logrus.DebugLevel, "healthz")
	c.JSON(http.StatusOK, "ok")
}
