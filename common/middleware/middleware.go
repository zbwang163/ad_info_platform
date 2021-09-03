package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zbwang163/ad_common/convert"
	"github.com/zbwang163/ad_common/time_utils"
	"github.com/zbwang163/ad_info_platform/common/consts"
	"github.com/zbwang163/ad_info_platform/common/utils"
	"github.com/zbwang163/ad_info_platform/common/utils/session"
	"os"
	"time"
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

	file, err := os.OpenFile(fmt.Sprintf("/var/log/ad_platform_info/%v.log", time_utils.Time20060102_15(time.Now())),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

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
