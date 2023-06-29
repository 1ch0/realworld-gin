package main

import (
	"github.com/1ch0/gin-template/cmd/server/app"
	"github.com/1ch0/gin-template/pkg/server/utils/log"
)

func main() {
	cmd := app.NewAPIServerCommand()
	if err := cmd.Execute(); err != nil {
		log.Logger.Fatalln(err)
	}
}
