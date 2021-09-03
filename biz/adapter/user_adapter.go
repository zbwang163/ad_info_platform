package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/zbwang163/ad_common/biz_error"
	contentRpc "github.com/zbwang163/ad_content_overpass"
	"github.com/zbwang163/ad_info_platform/biz/adapter/query"
	"github.com/zbwang163/ad_info_platform/biz/service/user"
	"github.com/zbwang163/ad_info_platform/common/clients"
	"github.com/zbwang163/ad_info_platform/common/utils"
)

type UserAdapter struct {
	userService *user.ServiceOfUser
}

func NewUserAdapter() *UserAdapter {
	return &UserAdapter{
		userService: user.NewServiceOfUser(),
	}
}

func (a UserAdapter) GetUserInfo(c *gin.Context) (utils.DTO, *biz_error.BizError) {
	params := &query.GetUserInfoQuery{}
	bizError := params.BindParams(c)
	if bizError != nil {
		return nil, bizError
	}
	return a.userService.GetUserDtoByCoreUserId(c, params.CoreUserId)
}

func (a UserAdapter) Login(c *gin.Context) (utils.DTO, *biz_error.BizError) {
	params := &query.LoginParams{}
	bizError := params.BindParams(c)
	if bizError != nil {
		return nil, bizError
	}
	resp, err := clients.ContentClient.Search(c, &contentRpc.SearchRequest{})
	if err != nil {
		return nil, biz_error.NewResourceError(err)
	}
	return resp, nil
}
