package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zbwang163/ad_common/biz_error"
	"github.com/zbwang163/ad_info_platform/biz/service/user/entity/dto"
	"github.com/zbwang163/ad_info_platform/biz/service/user/factory"
	"github.com/zbwang163/ad_info_platform/biz/service/user/repo"
)

type ServiceOfUser struct {
	adUserInfoRepo *repo.AdUserInfoRepo
	adCodeRefRepo  *repo.AdCodeRefRepo
	userDtoFactory *factory.UserDtoFactory
}

func NewServiceOfUser() *ServiceOfUser {
	return &ServiceOfUser{
		adUserInfoRepo: repo.NewAdUserInfoRepo(),
		adCodeRefRepo:  repo.NewAdCodeRefRepo(),
		userDtoFactory: factory.NewUserDtoFactory(),
	}
}

func (s ServiceOfUser) GetUserDtoByCoreUserId(c *gin.Context, coreUserId int64) (*dto.UserDto, *biz_error.BizError) {
	if coreUserId == 0 {
		return nil, nil
	}
	userPo, bizErr := s.adUserInfoRepo.QueryOne(c, map[string]interface{}{"core_user_id": coreUserId})
	if bizErr != nil {
		return nil, bizErr
	}
	codeMap, _ := s.getUserWorkCodeMap(c)
	return s.userDtoFactory.BuildUserDTOByPo(userPo, codeMap), nil
}

func (s ServiceOfUser) getUserWorkCodeMap(c *gin.Context) (map[string]string, *biz_error.BizError) {
	codePos, bizError := s.adCodeRefRepo.Query(c, map[string]interface{}{"type": 1})
	if bizError != nil {
		return nil, bizError
	}
	result := make(map[string]string)
	for _, po := range codePos {
		result[po.Code] = po.Value
	}
	return result, nil
}

//
//func (s ServiceOfUser) Login(c *gin.Context, userName, password, capture string)(*dto.LoginDto, *biz_error.BizError){
//
//	userPo,bizError := s.adUserInfoRepo.QueryOne(c, map[string]interface{}{"nickname":userName,"password":password})
//}
