package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"my_codes/ad_platform_info/common/consts"
	"my_codes/ad_platform_info/common/utils"
	"my_codes/ad_platform_info/common/utils/convert"
	"my_codes/ad_platform_info/common/utils/session"
)

func UserInfoMiddleware(c *gin.Context) {
	deviceIdStr := c.GetHeader("device_id")
	deviceId, _ := convert.StringToInt64(deviceIdStr)
	c.Set(consts.CtxDeviceId, deviceId)

	sessionId, _ := c.Cookie("session_id")
	c.Set(consts.CtxUserId, session.GetCoreUserIdFromSession(sessionId))
	c.Next()
}

func LogIdMiddleware(c *gin.Context) {
	logId := utils.GenerateLogId()
	c.Set(consts.LogId, logId)
	c.Next()
}

func LoggerMiddleware(c *gin.Context) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{}) //log日志序列化为json
	log.SetReportCaller(true)                 // 打印日志位置
	ctxLog := log.WithFields(logrus.Fields{
		consts.LogId: utils.GetCtxLogId(c),
		consts.Ip:    utils.GetLocalIp(),
	})
	c.Set(consts.Logger, ctxLog)
	c.Next()
}

func ResponseMiddleware(c *gin.Context) {
	c.Header(consts.LogId, utils.GetCtxLogId(c))
	c.Next()
}
