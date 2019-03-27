package conf

import "time"

const UserPrizeMax = 100000             // 用户每天最多抽奖次数
const IpPrizeMax = 50000             	// 同一个IP每天最多抽奖次数
const IpLimitMax = 500000               // 同一个IP每天最多抽奖次数

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