package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Register(r)
	r.Run() //8080端口
}
