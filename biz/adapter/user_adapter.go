package adapter

import (
	"github.com/gin-gonic/gin"
	"my_codes/ad_platform_info/biz/adapter/query"
	"my_codes/ad_platform_info/biz/service/user"
	user_dto "my_codes/ad_platform_info/biz/service/user/entity/dto"
	"my_codes/ad_platform_info/common/biz_error"
)

type UserAdapter struct {
	userService *user.ServiceOfUser
}

func NewUserAdapter() *UserAdapter {
	return &UserAdapter{
		userService: user.NewServiceOfUser(),
	}
}

func (a UserAdapter) GetUserInfo(c *gin.Context) (interface{}, *biz_error.BizError) {
	var result *user_dto.UserDto
	param, bizError := query.BindGetUserInfoQuery(c)
	if param == nil {
		return result, bizError
	}
	return a.userService.GetUserDtoByCoreUserId(c, param.CoreUserId)
}
