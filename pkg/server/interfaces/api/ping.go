package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pingApi struct {
	*gin.Engine
}

func NewPingApi(app *gin.Engine) Interface {
	return &pingApi{app}
}

func (h *pingApi) Register() {
	api := h.Group("/ping")
	api.GET("", h.healthz)
}

func (h *pingApi) healthz(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
