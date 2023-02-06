package internal

import (
	"douyin-simple/internal/conf"
	"fmt"
	"time"

	"github.com/google/wire"
	// "github.com/gorilla/securecookie"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	// "github.com/kataras/iris/v12/websocket"
)

var ProviderSet = wire.NewSet(New)

// type Configurator func(*Bootstrapper)

// bootstrap 与 server 的区别：
// bootstrap 负责与业务无关的配置，比如应用名称，应用所有者
// server 只负责业务内的配置，
// 比如静态资源路径，中间件(是否需要 logger，是否需要 recover)等
type Bootstrapper struct {
	*iris.Application
	addr         string
	appName      string
	appOwner     string
	appSpawnDate time.Time

	Sessions *sessions.Sessions
}

// New返回一个新的Bootstrapper
func New(conf *conf.Bootstrap, server *iris.Application) *Bootstrapper {
	b := &Bootstrapper{
		addr:         conf.Addr,
		Application:  server,
		appName:      conf.AppName,
		appOwner:     conf.AppOwner,
		appSpawnDate: time.Now(),
	}

	return b
}

// Listen 启动http服务器
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) error {
	if addr != "" {
		fmt.Println("--------------", addr, "---------------")
		b.addr = addr
	}
	return b.Run(iris.Addr(b.addr), cfgs...)
}

// Configure接受配置并在Bootstraper的上下文中运行它们
// func (b *Bootstrapper) Configure(cs ...Configurator) {
// 	for _, c := range cs {
// 		c(b)
// 	}
// }
