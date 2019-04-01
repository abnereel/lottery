package routes

import (
	"github.com/abnereel/lottery/bootstrap"
	"github.com/abnereel/lottery/services"
	"github.com/abnereel/lottery/web/controllers"
	"github.com/kataras/iris/_examples/mvc/login/web/middleware"
	"github.com/kataras/iris/mvc"
)

func Congifure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	blackipService := services.NewBlackipService()
	userdayService := services.NewUserdayService()

	// 前端首页
	index := mvc.New(b.Party("/"))
	index.Register(userService, giftService, codeService, resultService, blackipService, userdayService)
	index.Handle(new(controllers.IndexController))

	// 后台首页
	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(userService, giftService, codeService, resultService, blackipService, userdayService)
	admin.Handle(new(controllers.AdminController))

	// 奖品管理
	adminGift := admin.Party("/gift")
	adminGift.Register(giftService)
	adminGift.Handle(new(controllers.AdminGiftController))

	// 优惠券管理
	adminCode := admin.Party("/code")
	adminCode.Register(codeService)
	adminCode.Handle(new(controllers.AdminCodeController))

	// 中奖记录管理
	adminResult := admin.Party("/result")
	adminResult.Register(resultService)
	adminResult.Handle(new(controllers.AdminResultController))

	// 用户管理
	adminUser := admin.Party("/user")
	adminUser.Register(userService)
	adminUser.Handle(new(controllers.AdminUserController))

	// IP黑名单管理
	adminBlackip := admin.Party("/blackip")
	adminBlackip.Register(blackipService)
	adminBlackip.Handle(new(controllers.AdminBlackipController))

	// RPC
	rpc := mvc.New(b.Party("/rpc"))
	rpc.Register(userService, giftService, codeService, resultService, userdayService, blackipService)
	rpc.Handle(new(controllers.RpcController))
}
