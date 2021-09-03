package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/zbwang163/ad_common/biz_error"
	"github.com/zbwang163/ad_common/env"
	"github.com/zbwang163/ad_info_platform/common/consts"

	"net"
	"net/http"
	"time"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type DTO interface{}
type Handler func(*gin.Context) (DTO, *biz_error.BizError)

func HandlerFunc(f Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, bizError := f(c)

		if bizError == nil || bizError.Code == 0 {
			c.JSON(http.StatusOK, Response{0, "success", data})
		} else {
			if env.IsDev() {
				c.JSON(http.StatusOK, Response{bizError.Code, bizError.Message, bizError.ErrorMsg})
			} else {
				c.JSON(http.StatusOK, Response{bizError.Code, bizError.Message, ""})
			}
		}
	}
}

// GetLocalIp 获取本机的io
func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GenerateLogId 生成log id
func GenerateLogId() string {
	t := time.Now().Format("200601021504")
	bytes, _ := uuid.GenerateRandomBytes(11)
	return fmt.Sprintf("%v%x", t, bytes)
}

// GetCtxLogId 从context中获取log id
func GetCtxLogId(c *gin.Context) string {
	if logId, ok := c.Value(consts.LogId).(string); ok {
		return logId
	} else {
		return ""
	}
}
