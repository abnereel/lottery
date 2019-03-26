package controllers

import (
	"fmt"
	"github.com/abnereel/lottery/comm"
	"github.com/abnereel/lottery/models"
	"github.com/abnereel/lottery/services"
	"github.com/kataras/iris"
)

type IndexController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceBlackip services.BlackipService
	ServiceUserday services.UserdayService
}

// 首页
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "weclome to Go抽奖系统, <a href='/public/index.htm'>开始抽奖</a>"
}

// 获取奖品
func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	datalist := c.ServiceGift.GetAll()
	list := make([]models.LtGift, 0)
	for _, data := range datalist {
		if data.SysStatus == 0 {
			list = append(list, data)
		}
	}
	rs["gifts"] = list

	return rs
}

// 新的奖品
func (c *IndexController) GetNewprize() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	// TODO:

	return rs
}

// 登录
func (c *IndexController) GetLogin() {
	uid := comm.Random(100000)
	loginuser := models.ObjLoginuser{
		Uid:      uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index.htm?from=login")
}

// 退出
func (c *IndexController) GetLogout() {
	comm.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(), "/public/index.htm?from=logout")
}
