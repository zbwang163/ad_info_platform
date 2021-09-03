package dto

type UserDto struct {
	CoreUserId  int64    `json:"core_user_id"`
	Nickname    string   `json:"nickname"`
	AvatarUrl   string   `json:"avatar_url"`
	Description string   `json:"description"`
	Career      []string `json:"career"`
	Company     []string `json:"company"`
	Industry    []string `json:"industry"`
	Role        int64    `json:"role"`     //用户角色，0:普通用户；1:创意号；2:创作达人；
	RoleUrl     string   `json:"role_url"` // 角色的图标
}

var RoleUrlMap = map[int64]string{
	0: "",
	1: "",
	2: "",
}

type LoginDto struct {
	SessionId  string `json:"session_id"`
	CoreUserId int64  `json:"core_user_id"`
}
