package controllers

import (
	"github.com/abnereel/lottery/models"
	"time"
)

func (c *IndexController) checkBlackUser(uid int) (bool, *models.LtUser) {
	info := c.ServiceUser.Get(uid)
	if info != nil && info.Blacktime > int(time.Now().Unix()) {
		// 黑名单存在且有效
		return false, info
	}
	return true, info
}