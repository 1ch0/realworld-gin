package bcode

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

// Bcode business error code
type Bcode struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func ErrServer(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, fiber.Map{
		"code":    500,
		"message": "The service has lapsed.",
	})
}

// ErrForbidden check user perms failure
func ErrForbidden(ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, fiber.Map{
		"code":    403,
		"message": "403 Forbidden",
	})
}

// ErrUnauthorized check user auth failure
func ErrUnauthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, fiber.Map{
		"code":    401,
		"message": "401 Unauthorized",
	})
}

// ErrNotFound the request resource is not found
func ErrNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, fiber.Map{
		"code":    404,
		"message": "404 Not Found",
	})
}

// ErrUpstreamNotFound the proxy upstream is not found
func ErrUpstreamNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusBadGateway, fiber.Map{
		"code":    502,
		"message": "Upstream not found",
	})
}
