package api

import (
	"github.com/gin-gonic/gin"
)

type Interface interface {
	Register()
}

var registeredAPI []Interface

// RegisterAPI Register API handler
func RegisterAPI(ws ...Interface) {
	for _, w := range ws {
		registeredAPI = append(registeredAPI, w)
	}
}

func GetRegisteredAPI() []Interface {
	return registeredAPI
}

func InitAPIBean(server *gin.Engine) []interface{} {
	RegisterAPI(NewPingApi(server), NewHealthzApi(server))

	var beans []interface{}
	for i := range registeredAPI {
		beans = append(beans, registeredAPI[i])
	}
	beans = append(beans)
	return beans
}
