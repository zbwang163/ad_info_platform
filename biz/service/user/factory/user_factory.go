package factory

import (
	"github.com/zbwang163/ad_info_platform/biz/service/user/entity/dto"
	"github.com/zbwang163/ad_info_platform/biz/service/user/entity/po"
	"github.com/zbwang163/ad_info_platform/common/consts"
	"strings"
)

type UserDtoFactory struct {
}

func NewUserDtoFactory() *UserDtoFactory {
	return &UserDtoFactory{}
}

func (UserDtoFactory) BuildUserDTOByPo(po *po.AdUserInfoPo, codeRef map[string]string) *dto.UserDto {
	if po == nil {
		return nil
	}
	careerCodes := strings.Split(po.Career, consts.DbSeparator)
	industryCodes := strings.Split(po.Industry, consts.DbSeparator)
	companyCodes := strings.Split(po.Company, consts.DbSeparator)

	var careers, industries, companies []string
	for _, code := range careerCodes {
		careers = append(careers, codeRef[code])
	}

	for _, code := range industryCodes {
		industries = append(industries, codeRef[code])
	}

	for _, code := range companyCodes {
		companies = append(companies, codeRef[code])
	}

	roleUrl := dto.RoleUrlMap[po.Role]

	return &dto.UserDto{
		CoreUserId:  po.CoreUserId,
		Nickname:    po.Nickname,
		AvatarUrl:   po.AvatarUrl,
		Description: po.Description,
		Career:      careers,
		Company:     companies,
		Industry:    industries,
		Role:        po.Role,
		RoleUrl:     roleUrl,
	}
}
