// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysUserThirdBind = "sys_user_third_bind"

// SysUserThirdBind mapped from table <sys_user_third_bind>
type SysUserThirdBind struct {
	ID          int64          `gorm:"column:id;type:int(255);primaryKey;autoIncrement:true;comment:主键" json:"id"`     // 主键
	UserID      int64          `gorm:"column:user_id;type:int(255);comment:用户ID" json:"userId"`                        // 用户ID
	LoginType   string         `gorm:"column:login_type;type:varchar(500);comment:三方登录类型" json:"loginType"`            // 三方登录类型
	Openid      string         `gorm:"column:openid;type:varchar(500);comment:三方唯一标识" json:"openid"`                   // 三方唯一标识
	AccessToken string         `gorm:"column:access_token;type:varchar(255);comment:三方Token" json:"accessToken"`       // 三方Token
	CreateDept  int64          `gorm:"column:create_dept;type:int(11);comment:创建部门" json:"createDept"`                 // 创建部门
	CreateBy    int64          `gorm:"column:create_by;type:int(11);comment:创建者" json:"createBy"`                      // 创建者
	CreateTime  time.Time      `gorm:"column:create_time;autoCreateTime;type:datetime;comment:创建时间" json:"createTime"` // 创建时间
	UpdateBy    int64          `gorm:"column:update_by;type:int(11);comment:更新者" json:"updateBy"`                      // 更新者
	UpdateTime  time.Time      `gorm:"column:update_time;autoUpdateTime;type:datetime;comment:更新时间" json:"updateTime"` // 更新时间
	DeleteTime  gorm.DeletedAt `gorm:"column:delete_time;type:datetime;comment:删除时间" json:"deleteTime"`                // 删除时间
}

// TableName SysUserThirdBind's table name
func (*SysUserThirdBind) TableName() string {
	return TableNameSysUserThirdBind
}
