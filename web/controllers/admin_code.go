package controllers

import (
	"fmt"
	"github.com/abnereel/lottery/comm"
	"github.com/abnereel/lottery/conf"
	"github.com/abnereel/lottery/models"
	"github.com/abnereel/lottery/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strings"
)

type AdminCodeController struct {
	Ctx            iris.Context
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceBlackip services.BlackipService
	ServiceUserday services.UserdayService
}

func (c *AdminCodeController) Get() mvc.Result {
	giftId := c.Ctx.URLParamIntDefault("gift_id", 0)
	page := c.Ctx.URLParamIntDefault("page", 1)
	size := 100
	pagePrev := ""
	pageNext := ""
	// 数据列表
	var datalist []models.LtCode
	var total int = 0
	// TODO:
	if giftId > 0 {
		datalist = c.ServiceCode.Search(giftId)
	} else {
		datalist = c.ServiceCode.GetAll(page, size)
	}
	total = (page-1)*size + len(datalist)
	if len(datalist) >= size {
		if giftId > 0 {
			total = int(c.ServiceCode.CountByGift(giftId))
		} else {
			total = int(c.ServiceCode.CountAll())
		}
		pageNext = fmt.Sprintf("%d", page+1)
	}
	if page > 1 {
		pagePrev = fmt.Sprintf("%d", page-1)
	}
	return mvc.View{
		Name: "admin/code.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "code",
			"GiftId":   giftId,
			"Datalist": datalist,
			"Total":    total,
			"PagePrev": pagePrev,
			"PageNext": pageNext,
			/*"CodeNum":  num,
			"CacheNum": cacheNum,*/
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminCodeController) PostImport() {
	giftId := c.Ctx.URLParamIntDefault("gift_id", 0)
	if giftId < 1 {
		_, _ = c.Ctx.Text("没有指定奖品ID，无法进行导入，<a href='' onclick='history.go(-1);return false;'>返回</a>")
		return
	}
	gift := c.ServiceGift.Get(giftId)
	if gift == nil || gift.Id < 1 || gift.Gtype != conf.GtypeCodeDiff {
		_, _ = c.Ctx.Text("奖品信息不存在或者奖品类型不是差异化优惠券，无法导入，<a href='' onclick='history.go(-1);return false;'>返回</a>")
		return
	}
	codes := c.Ctx.PostValue("codes")
	now := comm.NowUnix()
	list := strings.Split(codes, "\n")
	sucNum := 0
	errNum := 0
	for _, code := range list {
		code := strings.TrimSpace(code)
		if code != "" {
			data := &models.LtCode{
				GiftId: giftId,
				Code: code,
				SysCreated: now,
			}
			err := c.ServiceCode.Create(data)
			if err != nil {
				errNum++
			} else {
				sucNum++
				// TODO: 成功导入数据库，下一步还需要导入缓存
			}
		}
	}
	_, _ = c.Ctx.HTML(fmt.Sprintf("成功导入%d条，导入失败%d条，<a href='/admin/code?gift_id=%d' >返回</a>", sucNum, errNum, giftId))
}

func (c *AdminCodeController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		_ = c.ServiceCode.Delete(id)
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/code"
	}
	return mvc.Response{
		Path: refer,
	}
}

func (c *AdminCodeController) GetReset() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		_ = c.ServiceCode.Update(&models.LtCode{Id: id, SysStatus: 0}, []string{"sys_status"})
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/code"
	}
	return mvc.Response{
		Path: refer,
	}
}
