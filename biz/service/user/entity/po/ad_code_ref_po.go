package po

import (
	"gorm.io/gorm"
)

type AdCodeRefPo struct {
	gorm.Model
	Code  string
	Value string
	Type  int64
	Extra string
}

func (*AdCodeRefPo) TableName() string {
	return "ad_code_ref"
}
