package query

import (
	"errors"
	"github.com/gin-gonic/gin"
	"my_codes/ad_platform_info/common/biz_error"
	"my_codes/ad_platform_info/common/consts"
	"my_codes/ad_platform_info/common/utils/convert"
)

type GetUserInfoParams struct {
	CoreUserId string `json:"core_user_id" binding:"required""`
}

type GetUserInfoQuery struct {
	CoreUserId int64 `json:"core_user_id"`
}

func BindGetUserInfoQuery(c *gin.Context) (*GetUserInfoQuery, *biz_error.BizError) {
	var params GetUserInfoParams
	err := c.ShouldBind(&params)
	if err != nil {
		return nil, biz_error.NewParamError(err)
	}
	coreUserId, _ := convert.StringToInt64(params.CoreUserId)
	ctxUserId := c.Value(consts.CtxUserId).(int64)
	if coreUserId != ctxUserId {
		return nil, biz_error.NewResourceError(errors.New("uid from param and context not equal"))
	}
	return &GetUserInfoQuery{CoreUserId: coreUserId}, nil
}
