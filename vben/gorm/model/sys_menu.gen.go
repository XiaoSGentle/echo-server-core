// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"github.com/super-sunshines/echo-server-core/core"
	"gorm.io/gorm"
)

const TableNameSysMenu = "sys_menu"

// SysMenu mapped from table <sys_menu>
type SysMenu struct {
	ID             int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true;comment:id" json:"id"`             // id
	Pid            int64          `gorm:"column:pid;type:int(11);comment:父目录" json:"pid"`                                        // 父目录
	Type           int64          `gorm:"column:type;type:int(1) unsigned zerofill;not null;comment:目录类型 0:目录 1:接口" json:"type"` // 目录类型 0:目录 1:接口
	APICode        string         `gorm:"column:api_code;type:varchar(255);comment:接口代码" json:"apiCode"`                         // 接口代码
	APIDescription string         `gorm:"column:api_description;type:varchar(255);comment:接口描述" json:"apiDescription"`           // 接口描述
	MetaID         int64          `gorm:"column:meta_id;type:int(11);comment:metaID" json:"metaId"`                              // metaID
	Name           string         `gorm:"column:name;type:varchar(255);comment:路由名称" json:"name"`                                // 路由名称
	Path           string         `gorm:"column:path;type:varchar(255);comment:访问路径" json:"path"`                                // 访问路径
	Component      string         `gorm:"column:component;type:varchar(255);comment:组件地址" json:"component"`                      // 组件地址
	CreateDept     int64          `gorm:"column:create_dept;type:int(11);comment:创建部门" json:"createDept"`                        // 创建部门
	CreateBy       int64          `gorm:"column:create_by;type:int(11);comment:创建者" json:"createBy"`                             // 创建者
	CreateTime     core.Time      `gorm:"column:create_time;autoCreateTime;type:datetime;comment:创建时间" json:"createTime"`        // 创建时间
	UpdateBy       int64          `gorm:"column:update_by;type:int(11);comment:更新者" json:"updateBy"`                             // 更新者
	UpdateTime     core.Time      `gorm:"column:update_time;autoUpdateTime;type:datetime;comment:更新时间" json:"updateTime"`        // 更新时间
	DeleteTime     gorm.DeletedAt `gorm:"column:delete_time;type:datetime;comment:删除时间" json:"deleteTime"`                       // 删除时间
}

// TableName SysMenu's table name
func (*SysMenu) TableName() string {
	return TableNameSysMenu
}
