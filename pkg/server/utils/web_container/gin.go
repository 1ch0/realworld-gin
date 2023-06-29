package web_container

import "github.com/gin-gonic/gin"

type GinWebContainer struct {
	*gin.Engine
}

func NewGinEngine() WebContainer {
	return &GinWebContainer{gin.New()}
}

func (g *GinWebContainer) RunWebContainer(addr string) error {
	return g.Run(addr)
}
