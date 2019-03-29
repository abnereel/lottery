package cron

import (
	"github.com/abnereel/lottery/comm"
	"github.com/abnereel/lottery/services"
	"github.com/abnereel/lottery/web/utils"
	"log"
	"time"
)

func ConfigureAppOneCron() {
	go resetAllGiftPrizeData()
	go distributionAllGiftPool()
}

func resetAllGiftPrizeData() {
	giftService := services.NewGiftService()
	nowTime := comm.NowUnix()
	// 从数据库读取会更新缓存
	list := giftService.GetAll(false)
	for _, giftInfo := range list {
		if giftInfo.PrizeTime > 0 &&
			(giftInfo.PrizeData == "" || giftInfo.PrizeEnd < nowTime) {
			log.Println("crontab start utils.ResetGiftPrizeData giftInfo=", giftInfo)
			utils.ResetGiftPrizeData(&giftInfo, giftService)
			// 重新写入缓存
			giftService.GetAll(true)
			log.Println("crontab end utils.ResetGiftPrizeData giftInfo=", giftInfo)
		}
	}

	// 每5分钟执行一次
	time.AfterFunc(5*time.Minute, resetAllGiftPrizeData)
}

func distributionAllGiftPool()  {
	log.Println("crontab start utils.distributionAllGiftPool")
	num := utils.DistributionGiftPool()
	log.Println("crontab end utils.distributionAllGiftPool, num=", num)

	// 每分钟执行一次
	time.AfterFunc(time.Minute, distributionAllGiftPool)
}
