package server

import (
	"douyin-simple/internal/conf"
	"douyin-simple/internal/service"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

const (
	DEBUG = "debug"
)

func NewHTTPServer(
	conf *conf.Server,
	coreService *service.CoreService,
	interactionService *service.InteractionService,
	socializeService *service.SocializeService,
) *iris.Application {
	// server := iris.Default()
	server := iris.New()

	// 静态资源配置
	// b.Favicon(StaticAssets + Favicon)
	server.HandleDir(conf.StaticUrl, iris.Dir(conf.StaticDir))

	// 中间件
	// 日志
	logger.New()
	if conf.IsLogger {
		server.Use(logger.New())
	}
	// 程序恢复
	if conf.IsRecover {
		server.Use(recover.New())
		setupErrorHandlers(server)
	}

	router := server.Party("/douyin")
	mvc.Configure(router, func(app *mvc.Application) {
		app.Handle(coreService)
	})
	mvc.Configure(router, func(app *mvc.Application) {
		app.Handle(interactionService)
	})
	mvc.Configure(router, func(app *mvc.Application) {
		app.Handle(socializeService)
	})
	return server
}

// SetupErrorHandlers准备http错误处理程序
// `(context.StatusCodeNotSuccessful`,  which defaults to >=400 (but you can change it).
func setupErrorHandlers(server *iris.Application) {
	server.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"status_code": ctx.GetStatusCode(),
			"status":      http.StatusText(ctx.GetStatusCode()),
			"message":     ctx.Values().GetString("message"),
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

// SetupSessions初始化会话(可选)。
// func setupSessions(server *iris.Application, expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
// 	server.Sessions = sessions.New(sessions.Config{
// 		Cookie:   "SECRET_SESS_COOKIE",
// 		Expires:  expires,
// 		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
// 	})
// }

// SetupViews 加载模板
// func setupViews(server *iris.Application, viewsDir string) {
// 	server.RegisterView(iris.HTML(viewsDir, ".html").Layout("shared/layout.html"))
// }

// SetupWebsockets准备websocket服务器。
// func setupWebsockets(server *iris.Application, endpoint string, handler websocket.ConnHandler) {
// 	ws := websocket.New(websocket.DefaultGorillaUpgrader, handler)
// 	server.Get(endpoint, websocket.Handler(ws))
// }