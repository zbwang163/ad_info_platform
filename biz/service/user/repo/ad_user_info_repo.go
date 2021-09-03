package repo

import (
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"github.com/zbwang163/ad_common/biz_error"
	"github.com/zbwang163/ad_common/db"
	"github.com/zbwang163/ad_info_platform/biz/service/user/entity/po"
	"github.com/zbwang163/ad_info_platform/common/consts"
)

type AdUserInfoRepo struct {
}

func NewAdUserInfoRepo() *AdUserInfoRepo {
	return &AdUserInfoRepo{}
}

func (r AdUserInfoRepo) QueryOne(c *gin.Context, condition map[string]interface{}) (*po.AdUserInfoPo, *biz_error.BizError) {
	var result *po.AdUserInfoPo
	err := db.ReadDb[consts.PSM].WithContext(c).Where(condition).Find(&result).Error
	if err != nil {
		logs.WithContext(c).Error(err.Error())
		return nil, biz_error.NewMysqlError(err)
	}
	return result, nil
}
