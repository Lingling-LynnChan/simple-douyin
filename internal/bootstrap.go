package internal

import (
	"douyin-simple/internal/conf"
	"fmt"
	"time"

	"github.com/google/wire"
	"github.com/gorilla/securecookie"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/sessions"
	// "github.com/kataras/iris/v12/websocket"
)

var ProviderSet = wire.NewSet(New)

// type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	addr         string
	appName      string
	appOwner     string
	appSpawnDate time.Time

	Sessions *sessions.Sessions
}

// New返回一个新的Bootstrapper
func New(conf *conf.Server, server *iris.Application) *Bootstrapper {
	b := &Bootstrapper{
		addr:         conf.Addr,
		Application:  server,
		appName:      conf.AppName,
		appOwner:     conf.AppOwner,
		appSpawnDate: time.Now(),
	}

	// for _, cfg := range cfgs {
	// 	cfg(b)
	// }

	return b
}

// SetupSessions初始化会话(可选)。
func (b *Bootstrapper) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + b.appName,
		Expires:  expires,
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
	})
}

// SetupErrorHandlers准备http错误处理程序
// `(context.StatusCodeNotSuccessful`,  which defaults to >=400 (but you can change it).
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.appName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		if err := ctx.View("shared/error.html"); err != nil {
			ctx.HTML("<h3>%s</h3>", err.Error())
			return
		}
	})
}

const (
	// StaticAssets是公共资产的根目录，如图像，css, js。
	StaticAssets = "./public/"
	// Favicon是我们应用程序相对于“StaticAssets”的Favicon路径。
	Favicon = "favicon.ico"
)

// 引导程序准备应用程序
//
// Returns itself.
func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	// views
	// b.SetupViews("./views")

	// session
	b.SetupSessions(24*time.Hour,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	b.SetupErrorHandlers()

	// 静态资源配置
	// b.Favicon(StaticAssets + Favicon)
	b.HandleDir("/public", iris.Dir(StaticAssets))

	// 中间件
	b.Use(recover.New())
	b.Use(logger.New())

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

// SetupViews 加载模板
// func (b *Bootstrapper) SetupViews(viewsDir string) {
// 	b.RegisterView(iris.HTML(viewsDir, ".html").Layout("shared/layout.html"))
// }

// SetupWebsockets准备websocket服务器。
// func (b *Bootstrapper) SetupWebsockets(endpoint string, handler websocket.ConnHandler) {
// 	ws := websocket.New(websocket.DefaultGorillaUpgrader, handler)
// 	b.Get(endpoint, websocket.Handler(ws))
// }

// Configure接受配置并在Bootstraper的上下文中运行它们
// func (b *Bootstrapper) Configure(cs ...Configurator) {
// 	for _, c := range cs {
// 		c(b)
// 	}
// }
