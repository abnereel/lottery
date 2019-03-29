package conf

import "time"

const UserPrizeMax = 100000 // 用户每天最多抽奖次数
const IpPrizeMax = 50000    // 同一个IP每天最多抽奖次数
const IpLimitMax = 500000   // 同一个IP每天最多抽奖次数

var PrizeDataRandomDayTime = [100]int{
	// 24 * 3 = 72 平均3%的机会
	// 100 - 72 = 28 剩余28%的机会
	// 7 * 4 = 28 剩下的分别给到7个时段增加4%的机会
	0, 0, 0,
	1, 1, 1,
	2, 2, 2,
	3, 3, 3,
	4, 4, 4,
	5, 5, 5,
	6, 6, 6,
	7, 7, 7,
	8, 8, 8,
	9, 9, 9, 9, 9, 9, 9,
	10, 10, 10, 10, 10, 10, 10,
	11, 11, 11,
	12, 12, 12,
	13, 13, 13,
	14, 14, 14,
	15, 15, 15, 15, 15, 15, 15,
	16, 16, 16, 16, 16, 16, 16,
	17, 17, 17, 17, 17, 17, 17,
	18, 18, 18,
	19, 19, 19,
	20, 20, 20, 20, 20, 20, 20,
	21, 21, 21, 21, 21, 21, 21,
	22, 22, 22,
	23, 23, 23,
}

const SysTimeform = "2006-01-02 15:04:05"
const SysTimeformShort = "2006-01-02"

var SysTimeLocation, _ = time.LoadLocation("Asia/Chongqing")
var SignSecret = []byte("0123456789abcdef")
var CookieSecret = "hellolottery"

const GtypeVirtual = 0   // 虚拟币
const GtypeCodeSame = 1  // 虚拟券，相同的码
const GtypeCodeDiff = 2  // 虚拟券，不同的码
const GtypeGiftSmall = 3 // 实物小奖
const GtypeGiftLarge = 4 // 实物大奖
