package utils

import (
	"encoding/json"
	"fmt"
	"github.com/abnereel/lottery/comm"
	"github.com/abnereel/lottery/conf"
	"github.com/abnereel/lottery/datasource"
	"github.com/abnereel/lottery/models"
	"github.com/abnereel/lottery/services"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

// 重置一个奖品的发奖周期信息
// 奖品剩余数量也会重新设置为当前奖品数量
// 奖品的奖品池有效数量则会设置为空
// 奖品数量、发放周期等设置有修改的时候，也需要重置
//【难点】根据发奖周期，重新更新发奖计划
func ResetGiftPrizeData(giftInfo *models.LtGift, giftService services.GiftService) {
	if giftInfo == nil || giftInfo.Id < 1 {
		return
	}
	id := giftInfo.Id
	nowTime := comm.NowUnix()
	// 不能发奖，不需要设置发奖周期
	if giftInfo.SysStatus == 1 || giftInfo.TimeBegin >= nowTime ||
		giftInfo.TimeEnd <= nowTime || giftInfo.LeftNum <= 0 || giftInfo.PrizeNum <= 0 {
		if giftInfo.PrizeData != "" {
			// 清空旧的发奖计划数据
			clearGiftPrizeData(giftInfo, giftService)
		}
		return
	}
	//没有设置发奖周期
	dayNum := giftInfo.PrizeTime
	if dayNum <= 0 {
		setGiftPool(id, giftInfo.LeftNum)
		return
	}
	//重置发奖计划数据
	setGiftPool(id, 0)
	// 实际的奖品计划分布运算
	prizeNum := giftInfo.PrizeNum
	avgNum := prizeNum / dayNum
	// 每天可以分配的奖品数
	dayPrizeNum := make(map[int]int)
	if avgNum >= 1 {
		for day := 0; day < dayNum; day++ {
			dayPrizeNum[day] = avgNum
		}
	}
	// 剩下的随机分配到任意哪天
	prizeNum -= dayNum * avgNum
	for prizeNum > 0 {
		prizeNum--
		day := comm.Random(dayNum)
		_, ok := dayPrizeNum[day]
		if !ok {
			dayPrizeNum[day] = 1
		} else {
			dayPrizeNum[day] += 1
		}
	}
	// 每天的map，每小时的map，60分钟的数组，奖品数
	prizeData := make(map[int]map[int][60]int)
	for day, num := range dayPrizeNum {
		// 计算出来这一天的发奖计划
		dayPrizeData := getGiftPrizeDataOneDay(num)
		prizeData[day] = dayPrizeData
	}
	// 将周期内每天、每小时、每分钟的数据 prizeData格式化（[时间:数量]）
	datalist := formatGiftPrizeData(nowTime, dayNum, prizeData)
	str, err := json.Marshal(datalist)
	if err != nil {
		log.Println("prizedata.ResetGiftPrizeDAta json error=", err)
	} else {
		// 保存奖品的分布计划数据
		info := &models.LtGift{
			Id:         giftInfo.Id,
			PrizeNum:   giftInfo.PrizeNum,
			PrizeData:  string(str),
			PrizeBegin: nowTime,
			PrizeEnd:   nowTime + dayNum*86400,
			SysUpdated: nowTime,
		}
		err := giftService.Update(info, nil)
		if err != nil {
			log.Println("prizedata.ResetGiftPrizeData giftService.Update", info, ", error=", err)
		}
	}
}

// 获取当前奖品池中的奖品数量
func GetGiftPoolNum(id int) int {
	key := "gift_pool"
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("HGET", key, id)
	if err != nil {
		log.Println("prizedata.GetGiftPoolNum HGET error=", err)
		return 0
	} else {
		num := comm.GetInt64(rs, 0)
		return int(num)
	}
}

func PrizeGift(id, leftNum int) bool {
	ok := false
	ok = prizeServGift(id)
	if ok {
		giftService := services.NewGiftService()
		rows, err := giftService.DecrLeftNum(id, 1)
		if rows < 1 || err != nil {
			log.Println("prizedata.PrizeGift giftService.DecrLeftNum error=", err, ", rows=", rows)
			return false
		}
	}
	return ok
}

// 优惠券类的发放
func PrizeCodeDiff(id int, codeService services.CodeService) string {
	return prizeServCodeDiff(id, codeService)
}

// 导入新的优惠券编码
func ImportCacheCodes(id int, code string) bool {
	key := fmt.Sprintf("gift_code_%d", id)
	cacheObj := datasource.InstanceCache()
	_, err := cacheObj.Do("SADD", key, code)
	if err != nil {
		log.Println("prizedata.ImportCacheCodes SADD error=", err)
		return false
	} else {
		return true
	}
}

// 重新整理优惠券的编码到缓存中
func RecacheCodes(id int, codeService services.CodeService) (sucNum, errNum int) {
	list := codeService.Search(id)
	if list == nil || len(list) <= 0 {
		return 0, 0
	}
	key := fmt.Sprintf("gift_code_%d", id)
	cacheObj := datasource.InstanceCache()
	tmpKey := "tmp_" + key
	for _, data := range list {
		if data.SysStatus == 0 {
			code := data.Code
			_, err := cacheObj.Do("SADD", tmpKey, code)
			if err != nil {
				log.Println("prizedata.RecacheCodes SADD error=", err)
				errNum++
			} else {
				sucNum++
			}
		}
	}
	_, err := cacheObj.Do("RENAME", tmpKey, key)
	if err != nil {
		log.Println("prizedata.RecacheCodes RENAME error=", err)
	}
	return sucNum, errNum
}

// 获取当前的缓存中编码数量
// 返回，剩余编码数量，缓冲中编码数量
func GetCacheCodeNum(id int, codeService services.CodeService) (int, int) {
	num := 0
	cacheNum := 0
	//统计数据库中的
	list := codeService.Search(id)
	if len(list) > 0 {
		for _, data := range list {
			if data.SysStatus == 0 {
				num++
			}
		}
	}
	//统计redis缓存中的
	key := fmt.Sprintf("gift_code_%d", id)
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("SCARD", key)
	if err != nil {
		log.Println("prizedata.RecacheCodes SCARD error=", err)
	} else {
		cacheNum = int(comm.GetInt64(rs, 0))
	}
	return num, cacheNum
}

func prizeServGift(id int) bool {
	key := "gift_pool"
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("HINCRBY", key, id, -1)
	if err != nil {
		log.Println("prizedata.prizeServGift HINCRBY error=", err)
		return false
	}
	num := comm.GetInt64(rs, -1)
	if num >= 0 {
		return true
	} else {
		return false
	}
}

// 优惠券发放，使用redis的方式发放
func prizeServCodeDiff(id int, codeService services.CodeService) string {
	key := fmt.Sprintf("gift_code_%d", id)
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("SPOP", key)
	if err != nil {
		log.Println("prizedata.prizeServCodeDiff SPOP error=", err)
		return ""
	}
	code := comm.GetString(rs, "")
	if code == "" {
		log.Println("prizedata.prizeServCodeDiff rs=", rs)
		return ""
	}
	_ = codeService.UpdateByCode(&models.LtCode{
		Code:       code,
		SysStatus:  2,
		SysUpdated: comm.NowUnix(),
	}, nil)
	return code
}

func prizeLocalCodeDiff(id int, codeService services.CodeService) string {
	lockUid := 0 - id - 100000000
	LockLucky(lockUid)
	defer UnlockLucky(lockUid)

	codeId := 0
	codeInfo := codeService.NextUsingCode(id, codeId)
	if codeInfo != nil && codeInfo.Id > 0 {
		codeInfo.SysStatus = 2
		codeInfo.SysUpdated = comm.NowUnix()
		_ = codeService.Update(codeInfo, nil)
	} else {
		log.Println("prizedata.PrizeCodeDiff num codeInfo, gift_id=", id)
		return ""
	}
	return codeInfo.Code
}

// 清空旧的发奖计划数据
func clearGiftPrizeData(giftInfo *models.LtGift, giftService services.GiftService) {
	info := &models.LtGift{
		Id:        giftInfo.Id,
		PrizeData: "",
	}
	err := giftService.Update(info, []string{"prize_data"})
	if err != nil {
		log.Println("prizedata.clearGiftPrizeData giftService.Update", info, " error=", err)
	}
	setGiftPool(giftInfo.Id, 0)
}

// 设置奖品池的库存数量
func setGiftPool(id int, num int) {
	key := "gift_pool"
	cacheObj := datasource.InstanceCache()
	_, err := cacheObj.Do("HSET", key, id, num)
	if err != nil {
		log.Println("prizedata.setGiftPool error=", err)
	}
}

// 计算出一天的发奖计划
func getGiftPrizeDataOneDay(num int) map[int][60]int {
	rs := make(map[int][60]int)
	// 计算24小时各自的奖品数
	hourData := [24]int{}
	if num > 100 {
		hourNum := 0
		for _, h := range conf.PrizeDataRandomDayTime {
			hourData[h]++
		}
		for h := 0; h < 24; h++ {
			d := hourData[h]
			n := num * d / 100
			hourData[h] = n
			hourNum += n
		}
		num -= hourNum
	}
	// 剩余的随机分配
	for num > 0 {
		num--
		hourIndex := comm.Random(100)
		h := conf.PrizeDataRandomDayTime[hourIndex]
		hourData[h]++
	}
	// 将每个小时内的奖品数量分配到60分钟
	for h, hnum := range hourData {
		if hnum <= 0 {
			continue
		}
		minuteData := [60]int{}
		if hnum >= 60 {
			avgMinute := hnum / 60
			for i := 0; i < 60; i++ {
				minuteData[i] = avgMinute
			}
			hnum -= avgMinute * 60
		}
		// 剩余的随机分配
		for hnum > 0 {
			hnum--
			m := comm.Random(60)
			minuteData[m]++
		}
		rs[h] = minuteData
	}
	return rs
}

// 将每天、每小时、每分钟的奖品数量，格式化成具体到一个时间（分钟）的奖品数量
// 结构为： [day][hour][minute]num
// Result: [][时间:数量]
func formatGiftPrizeData(nowTime, dayNum int, prizeData map[int]map[int][60]int) [][2]int {
	rs := make([][2]int, 0)
	nowHour := time.Now().Hour()
	// 处理日期的数据
	for dn := 0; dn < dayNum; dn++ {
		dayData, ok := prizeData[dn]
		if !ok {
			continue
		}
		dayTime := nowTime + dn*86400
		// 处理小时的数据
		for hn := 0; hn < 24; hn++ {
			hourData, ok := dayData[(hn+nowHour)%24]
			if !ok {
				continue
			}
			hourTime := dayTime + hn*3600
			// 处理分钟的数据
			for mn := 0; mn < 60; mn++ {
				num := hourData[mn]
				if num <= 0 {
					continue
				}
				minuteTime := hourTime + mn*60
				rs = append(rs, [2]int{minuteTime, num})
			}
		}
	}
	return rs
}

/**
 * 根据奖品的发奖计划，把设定的奖品数量放入奖品池
 * 需要每分钟执行一次
 *【难点】定时程序，根据奖品设置的数据，更新奖品池的数据
 */
func DistributionGiftPool() int {
	totalNum := 0
	now := comm.NowUnix()
	giftService := services.NewGiftService()
	list := giftService.GetAll(false)
	if list != nil && len(list) > 0 {
		for _, gift := range list {
			// 是否正常状态
			if gift.SysStatus != 0 {
				continue
			}
			// 是否限量产品
			if gift.PrizeNum < 1 {
				continue
			}
			// 时间段是否正常
			if gift.TimeBegin > now || gift.TimeEnd < now {
				continue
			}
			// 计划数据的长度太短，不需要解析和执行
			// 发奖计划，[[时间1,数量1],[时间2,数量2]]
			if len(gift.PrizeData) <= 7 {
				continue
			}
			var cronData [][2]int
			err := json.Unmarshal([]byte(gift.PrizeData), &cronData)
			if err != nil {
				log.Println("prizedata.DistributionGiftPool Unmarshal error", err)
			} else {
				index := 0
				giftNum := 0
				for i, data := range cronData {
					ct := data[0]
					num := data[1]
					if ct <= now {
						giftNum += num
						index = i + 1
					} else {
						break
					}
				}
				// 更新奖品池
				if giftNum > 0 {
					incrGiftPool(gift.Id, giftNum)
					totalNum += giftNum
				}
				// 更新奖品的发奖计划
				if index > 0 {
					if index >= len(cronData) {
						cronData = make([][2]int, 0)
					} else {
						cronData = cronData[index:]
					}
					str, err := json.Marshal(cronData)
					if err != nil {
						log.Println("prizedata.DistributionGiftPool Marshal(cronData)", cronData, "error=", err)
					}
					columns := []string{"prize_data"}
					err = giftService.Update(&models.LtGift{
						Id:        gift.Id,
						PrizeData: string(str),
					}, columns)
					if err != nil {
						log.Println("prizedata.DistributionGiftPool giftService.Update error=", err)
					}
				}
			}
		}
		if totalNum > 0 {
			giftService.GetAll(true)
		}
	}
	return totalNum
}

// 往奖品池增加奖品数量，redis缓存，根据计划数据
func incrGiftPool(id, num int) int {
	key := "gift_pool"
	cacheObj := datasource.InstanceCache()
	rtNum, err := redis.Int64(cacheObj.Do("HINCRBY", key, id, num))
	if err != nil {
		log.Println("prizedata.incrGiftPool error=", err)
		return 0
	}
	if int(rtNum) < num {
		// 递增少于预期值，补偿一次
		num2 := num - int(rtNum)
		rtNum, err = redis.Int64(cacheObj.Do("HINCRBY", key, id, num2))
		if err != nil {
			log.Println("prizedata.incrGiftPool2 error=", err)
			return 0
		}
	}
	return int(rtNum)
}