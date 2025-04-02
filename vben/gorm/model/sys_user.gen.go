// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"echo-server-core/core"
	"time"

	"gorm.io/gorm"
)

const TableNameSysUser = "sys_user"

// SysUser mapped from table <sys_user>
type SysUser struct {
	ID             int64              `gorm:"column:id;type:int(255);primaryKey;autoIncrement:true;comment:主键" json:"id"`     // 主键
	Username       string             `gorm:"column:username;type:varchar(255);comment:用户名" json:"username"`                  // 用户名
	Password       string             `gorm:"column:password;type:varchar(255);comment:密码" json:"password"`                   // 密码
	NickName       string             `gorm:"column:nick_name;type:varchar(255);comment:昵称" json:"nickName"`                  // 昵称
	RealName       string             `gorm:"column:real_name;type:varchar(255);comment:真实姓名" json:"realName"`                // 真实姓名
	RoleCodeList   core.Array[string] `gorm:"column:role_code_list;type:json;comment:角色CODE列表" json:"roleCodeList"`           // 角色CODE列表
	Email          string             `gorm:"column:email;type:varchar(255);comment:邮箱地址" json:"email"`                       // 邮箱地址
	Avatar         string             `gorm:"column:avatar;type:varchar(255);comment:头像" json:"avatar"`                       // 头像
	LoginFailCount int64              `gorm:"column:login_fail_count;type:int(11);comment:登录失败次数" json:"loginFailCount"`      // 登录失败次数
	Phone          string             `gorm:"column:phone;type:varchar(11);comment:手机号" json:"phone"`                         // 手机号
	Status         int64              `gorm:"column:status;type:int(11);comment:状态" json:"status"`                            // 状态
	LastOnline     int64              `gorm:"column:last_online;type:bigint(20);comment:上次在线时间" json:"lastOnline"`            // 上次在线时间
	DepartmentID   int64              `gorm:"column:department_id;type:int(11);comment:部门ID" json:"departmentId"`             // 部门ID
	CreateDept     int64              `gorm:"column:create_dept;type:int(11);comment:创建部门" json:"createDept"`                 // 创建部门
	CreateBy       int64              `gorm:"column:create_by;type:int(11);comment:创建者" json:"createBy"`                      // 创建者
	CreateTime     time.Time          `gorm:"column:create_time;autoCreateTime;type:datetime;comment:创建时间" json:"createTime"` // 创建时间
	UpdateBy       int64              `gorm:"column:update_by;type:int(11);comment:更新者" json:"updateBy"`                      // 更新者
	UpdateTime     time.Time          `gorm:"column:update_time;autoUpdateTime;type:datetime;comment:更新时间" json:"updateTime"` // 更新时间
	DeleteTime     gorm.DeletedAt     `gorm:"column:delete_time;type:datetime;comment:删除时间" json:"deleteTime"`                // 删除时间
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}
