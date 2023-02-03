package server

import (
	"douyin-simple/internal/conf"
	"douyin-simple/internal/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func NewHTTPServer(conf *conf.Server, coreService *service.CoreService) *iris.Application {
	// todo
	// server := iris.New()
	server := iris.Default()
	router := server.Party("/douyin")
	mvc.Configure(router, func(app *mvc.Application) {
		app.Handle(coreService)
	})
	return server
}
