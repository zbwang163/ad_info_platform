package repo

import (
	"github.com/gin-gonic/gin"
	"my_codes/ad_platform_info/biz/service/user/entity/po"
	"my_codes/ad_platform_info/common/biz_error"
	"my_codes/ad_platform_info/common/clients"
	"my_codes/ad_platform_info/common/utils/logs"
)

type AdUserInfoRepo struct {
}

func NewAdUserInfoRepo() *AdUserInfoRepo {
	return &AdUserInfoRepo{}
}

func (r AdUserInfoRepo) QueryOne(c *gin.Context, condition map[string]interface{}) (*po.AdUserInfoPo, *biz_error.BizError) {
	var result *po.AdUserInfoPo
	err := clients.ReadDb.WithContext(c).Where(condition).Find(&result).Error
	if err != nil {
		logs.WithContext(c).Error(err.Error())
		return nil, biz_error.NewMysqlError(err)
	}
	return result, nil
}
