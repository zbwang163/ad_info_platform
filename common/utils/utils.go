package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	error2 "my_codes/ad_platform_info/common/biz_error"
	"my_codes/ad_platform_info/common/consts"
	"my_codes/ad_platform_info/common/utils/env"
	"net"
	"net/http"
	"time"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type Handler func(*gin.Context) (interface{}, *error2.BizError)

func HandlerFunc(f Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, bizError := f(c)

		if bizError == nil || bizError.Code == 0 {
			c.JSON(http.StatusOK, Response{0, "success", data})
		} else {
			if env.IsDev() {
				c.JSON(http.StatusOK, Response{bizError.Code, bizError.Message, bizError.Error})
			} else {
				c.JSON(http.StatusOK, Response{bizError.Code, bizError.Message, ""})
			}
		}
	}
}

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

func GenerateLogId() string {
	t := time.Now().Format("200601021504")
	bytes, _ := uuid.GenerateRandomBytes(11)
	return fmt.Sprintf("%v%x", t, bytes)
}

func GetCtxLogId(c *gin.Context) string {
	return c.Value(consts.LogId).(string)
}
