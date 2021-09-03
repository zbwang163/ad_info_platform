package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zbwang163/ad_info_platform/biz/adapter"
	"github.com/zbwang163/ad_info_platform/common/middleware"
	"github.com/zbwang163/ad_info_platform/common/utils"
	"net/http"
)

type Adapter struct {
	userAdapter *adapter.UserAdapter
}

func NewAdapter() *Adapter {
	return &Adapter{
		userAdapter: adapter.NewUserAdapter(),
	}
}

func Register(r *gin.Engine) {
	g := r.Group("/ad_info_platform")
	//g.Use(middleware.UserInfoMiddleware, middleware.LogIdMiddleware, middleware.LoggerMiddleware, middleware.ResponseMiddleware)
	g.Use(middleware.ResponseMiddleware)

	app := NewAdapter()

	userRouter := g.Group("/user")
	userRouter.POST("/info", utils.HandlerFunc(app.userAdapter.GetUserInfo))
	userRouter.POST("/login", utils.HandlerFunc(app.userAdapter.Login))
	userRouter.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"test": "wangzhibin"})
	})

}
