package main

import (
	"github.com/gin-gonic/gin"
	"my_codes/ad_platform_info/biz/adapter"
	"my_codes/ad_platform_info/common/middleware"
	"my_codes/ad_platform_info/common/utils"
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
	g := r.Group("/ad_platform_info")
	g.Use(middleware.UserInfoMiddleware, middleware.LogIdMiddleware, middleware.LoggerMiddleware, middleware.ResponseMiddleware)

	app := NewAdapter()

	userRouter := g.Group("/user")
	userRouter.POST("/info", utils.HandlerFunc(app.userAdapter.GetUserInfo))
}
