package controllers

import (
	"github.com/abnereel/lottery/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceBlackip services.BlackipService
	ServiceUserday services.UserdayService
}

func (c *AdminController) Get() mvc.Result {
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title": "管理后台",
			"Channel": "",
		},
		Layout: "admin/layout.html",
	}
}