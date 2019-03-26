package routes

import (
	"github.com/abnereel/lottery/bootstrap"
	"github.com/abnereel/lottery/services"
	"github.com/abnereel/lottery/web/controllers"
	"github.com/kataras/iris/mvc"
)

func Congifure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	blackipService := services.NewBlackService()
	userdayService := services.NewUserdayService()

	index := mvc.New(b.Party("/"))
	index.Register(
		userService,
		giftService,
		codeService,
		resultService,
		blackipService,
		userdayService,
	)
	index.Handle(new(controllers.IndexController))
}
