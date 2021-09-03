package repo

import (
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"github.com/zbwang163/ad_common/biz_error"
	"github.com/zbwang163/ad_common/db"
	"github.com/zbwang163/ad_info_platform/biz/service/user/entity/po"
	"github.com/zbwang163/ad_info_platform/common/consts"
)

type AdCodeRefRepo struct {
}

func NewAdCodeRefRepo() *AdCodeRefRepo {
	return &AdCodeRefRepo{}
}

func (r AdCodeRefRepo) QueryOne(c *gin.Context, condition map[string]interface{}) (*po.AdCodeRefPo, *biz_error.BizError) {
	var result *po.AdCodeRefPo
	err := db.ReadDb[consts.PSM].WithContext(c).Where(condition).First(&result).Error
	if err != nil {
		logs.WithContext(c).Error(err.Error())
		return nil, biz_error.NewMysqlError(err)
	}
	return result, nil
}

func (r AdCodeRefRepo) Query(c *gin.Context, condition map[string]interface{}) ([]*po.AdCodeRefPo, *biz_error.BizError) {
	var result []*po.AdCodeRefPo
	err := db.ReadDb[consts.PSM].WithContext(c).Where(condition).Find(&result).Error
	if err != nil {
		logs.WithContext(c).Error(err.Error())
		return nil, biz_error.NewMysqlError(err)
	}
	return result, nil
}
