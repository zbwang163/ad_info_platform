package session

import (
	"github.com/zbwang163/ad_common/convert"
	"github.com/zbwang163/ad_common/db"
	"github.com/zbwang163/ad_info_platform/common/consts"
)

func GetCoreUserIdFromSession(sessionId string) int64 {
	coreUserIdStr := db.Redis[consts.PSM].Get(sessionId).Val()
	coreUserId, _ := convert.StringToInt64(coreUserIdStr)
	return coreUserId
}
