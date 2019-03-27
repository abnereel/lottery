package utils

import (
	"fmt"
	"github.com/abnereel/lottery/datasource"
)

func getLuckyLockKey(uid int) string {
	return fmt.Sprintf("lucky_lock_%d", uid)
}

// 锁住用户
func LockLucky(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	rs, _ := cacheObj.Do("SET", key, 1, "EX", 3, "NX")
	if rs.(string) == "OK" {
		return true
	} else {
		return false
	}
}

// 对用户解锁
func UnlockLucky(uid int) bool {
	key := getLuckyLockKey(uid)
	cacheObj := datasource.InstanceCache()
	rs, _ := cacheObj.Do("DEL", key)
	if rs.(int64) > 0 {
		return true
	} else {
		return false
	}
}