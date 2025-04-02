package helper

import (
	"github.com/XiaoSGentle/echo-server-core/core"
	"github.com/XiaoSGentle/echo-server-core/vben/gorm/model"
)

func GenJwtByUserInfo(platform string, a model.SysUser) string {
	jwtStringModel, _ := core.GetTokenManager().GenJwtString(platform, core.ClaimsAdditions{
		UID:          a.ID,
		NickName:     a.NickName,
		Username:     a.Username,
		DepartmentId: a.DepartmentID,
		RoleCodes:    a.RoleCodeList,
	})
	return jwtStringModel
}
