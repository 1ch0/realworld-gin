package bcode

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UnHealthy the request resource is not found
func UnHealthy(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, Bcode{
		10000, "Unhealthy",
	})
}
