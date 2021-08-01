package test

import (
	"fmt"
	"my_codes/ad_platform_info/common/utils"
	"net"
	"testing"
)

func TestError(t *testing.T) {
	var err error
	t.Logf(err.Error())
}

func TestUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(utils.GenerateLogId())
	}
}

func TestIP(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}
