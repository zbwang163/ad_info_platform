package session

import (
	"my_codes/ad_platform_info/common/clients"
	"my_codes/ad_platform_info/common/utils/convert"
)

func GetCoreUserIdFromSession(sessionId string) int64 {
	clients.InitRedis()
	coreUserIdStr := clients.Redis.Get(sessionId).Val()
	coreUserId, _ := convert.StringToInt64(coreUserIdStr)
	return coreUserId
}
