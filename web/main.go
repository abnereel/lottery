package main

import (
	"fmt"
	"github.com/abnereel/lottery/bootstrap"
	"github.com/abnereel/lottery/web/middleware/identity"
	"github.com/abnereel/lottery/web/routes"
)

var port = 8080

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("Go抽奖系统", "Abner")
	app.Boostrap()
	app.Configure(identity.Configure, routes.Congifure)
	return app
}

func main() {
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
