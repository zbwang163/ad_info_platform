package po

import (
	"gorm.io/gorm"
)

type AdUserInfoPo struct {
	gorm.Model
	CoreUserId  int64  `gorm:"column:core_user_id,comment:'用户uid'"`
	Nickname    string `gorm:"column:nickname;comment:'昵称'"`
	AvatarUrl   string `gorm:"column:avatar_url;comment:'头像url'"`
	Description string `gorm:"column:description;comment:'个人描述'"`
	Career      string `gorm:"column:career;comment:'用户职业'"`
	Company     string `gorm:"column:company;comment:'用户公司'"`
	Industry    string `gorm:"column:industry;comment:'用户行业'"`
	Role        int64  `gorm:"column:role;comment:'用户角色，0:普通用户；1:创意号；2:创作达人；'"`
	Extra       string `gorm:"column:extra;comment:'额外字段'"`
}

func (*AdUserInfoPo) TableName() string {
	return "ad_user_info"
}
