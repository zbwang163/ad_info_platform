package query

import (
	"errors"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"github.com/zbwang163/ad_common/biz_error"
	"github.com/zbwang163/ad_common/convert"
	"github.com/zbwang163/ad_info_platform/common/consts"
)

type GetUserInfoQuery struct {
	CoreUserId int64 `json:"core_user_id" binding:"required"`
}

func (p *GetUserInfoQuery) BindParams(c *gin.Context) *biz_error.BizError {
	var params struct {
		CoreUserId string `json:"core_user_id" binding:"required"`
	}
	err := c.ShouldBind(&params)
	if err != nil {
		logs.WithContext(c).Warn(err)
		return biz_error.NewParamError(err)
	}
	coreUserId, _ := convert.StringToInt64(params.CoreUserId)
	ctxUserId := c.Value(consts.CtxUserId).(int64)
	if coreUserId != ctxUserId {
		return biz_error.NewResourceError(errors.New("uid from param and context not equal"))
	}
	p.CoreUserId = coreUserId
	return nil
}

type LoginParams struct {
	UserName string `json:"user_name" binding:"required,min=1,max=20"`
	Password string `json:"password" binding:"required,min=1,max=20"`
	Captcha  string `json:"captcha" binding:"required,len=4"`
}

func (p *LoginParams) BindParams(c *gin.Context) *biz_error.BizError {
	err := c.ShouldBind(p)
	if err != nil {
		return biz_error.NewParamError(err)
	}
	return nil
}
