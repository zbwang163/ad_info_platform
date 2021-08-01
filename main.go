package main

import (
	"github.com/gin-gonic/gin"
	"my_codes/ad_platform_info/common/clients"
)

func main() {
	r := gin.Default()
	Register(r)
	r.Run() //8080端口
}

func init() {
	clients.InitMysql()
	clients.InitRedis()
}
