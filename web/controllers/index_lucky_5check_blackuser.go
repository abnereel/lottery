package controllers

import (
	"github.com/abnereel/lottery/models"
	"github.com/abnereel/lottery/services"
	"time"
)

func (api *LuckyApi) checkBlackUser(uid int) (bool, *models.LtUser) {
	info := services.NewUserService().Get(uid)
	if info != nil && info.Blacktime > int(time.Now().Unix()) {
		// 黑名单存在且有效
		return false, info
	}
	return true, info
}