package repo

import (
	"github.com/gin-gonic/gin"
	"my_codes/ad_platform_info/biz/service/user/entity/po"
	"my_codes/ad_platform_info/common/biz_error"
	"my_codes/ad_platform_info/common/clients"
	"my_codes/ad_platform_info/common/utils/logs"
)

type AdCodeRefRepo struct {
}

func NewAdCodeRefRepo() *AdCodeRefRepo {
	return &AdCodeRefRepo{}
}

func (r AdCodeRefRepo) QueryOne(c *gin.Context, condition map[string]interface{}) (*po.AdCodeRefPo, *biz_error.BizError) {
	var result *po.AdCodeRefPo
	err := clients.ReadDb.WithContext(c).Where(condition).First(&result).Error
	if err != nil {
		logs.WithContext(c).Error(err.Error())
		return nil, biz_error.NewMysqlError(err)
	}
	return result, nil
}

func (r AdCodeRefRepo) Query(c *gin.Context, condition map[string]interface{}) ([]*po.AdCodeRefPo, *biz_error.BizError) {
	var result []*po.AdCodeRefPo
	err := clients.ReadDb.WithContext(c).Where(condition).Find(&result).Error
	if err != nil {
		logs.WithContext(c).Error(err.Error())
		return nil, biz_error.NewMysqlError(err)
	}
	return result, nil
}
