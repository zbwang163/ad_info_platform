package logs

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"my_codes/ad_platform_info/common/consts"
)

func WithContext(c *gin.Context) *logrus.Entry {
	if res, ok := c.Value(consts.Logger).(*logrus.Entry); ok {
		return res
	}
	return logrus.NewEntry(logrus.New())
}
