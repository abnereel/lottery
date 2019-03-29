package bootstrap

import (
	"github.com/abnereel/lottery/conf"
	"github.com/abnereel/lottery/cron"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"time"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application:  iris.New(),
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

//设置模板
func (b *Bootstrapper) SetupViews(viewDir string) {
	htmlEngine := iris.HTML(viewDir, ".html").Layout("shared/layout.html")
	htmlEngine.Reload(true)
	htmlEngine.AddFunc("FromUnixtimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeformShort)
	})
	htmlEngine.AddFunc("FromUnixtime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeform)
	})
	b.RegisterView(htmlEngine)
}

//设置异常处理
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}
		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			_, _ = ctx.JSON(err)
			return
		}
		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		_ = ctx.View("shared/error.html")
	})
}

//设置配置
func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, cfg := range cs {
		cfg(b)
	}
}

//计划任务
func (b *Bootstrapper) setupCron() {
	cron.ConfigureAppOneCron()
}

const (
	StaticAssets = "./public/"
	Favicon      = "favicon.ico"
)

func (b *Bootstrapper) Boostrap() *Bootstrapper {
	b.SetupViews("./views")
	b.SetupErrorHandlers()
	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	b.setupCron()

	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	_ = b.Run(iris.Addr(addr), cfgs...)
}
