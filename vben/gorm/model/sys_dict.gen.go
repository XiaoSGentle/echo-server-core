// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysDict = "sys_dict"

// SysDict mapped from table <sys_dict>
type SysDict struct {
	ID         int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	Module     int64          `gorm:"column:module;type:int(11)" json:"module"`
	Code       string         `gorm:"column:code;type:varchar(255)" json:"code"`
	Name       string         `gorm:"column:name;type:varchar(255)" json:"name"`
	ValueType  int64          `gorm:"column:value_type;type:int(11)" json:"valueType"`
	Describe   string         `gorm:"column:describe;type:varchar(255)" json:"describe"`
	Status     int64          `gorm:"column:status;type:int(11)" json:"status"`
	CreateDept int64          `gorm:"column:create_dept;type:int(11)" json:"createDept"`
	CreateBy   int64          `gorm:"column:create_by;type:int(11)" json:"createBy"`
	CreateTime time.Time      `gorm:"column:create_time;autoCreateTime;type:datetime" json:"createTime"`
	UpdateBy   int64          `gorm:"column:update_by;type:int(11)" json:"updateBy"`
	UpdateTime time.Time      `gorm:"column:update_time;autoUpdateTime;type:datetime" json:"updateTime"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:varchar(255)" json:"deleteTime"`
}

// TableName SysDict's table name
func (*SysDict) TableName() string {
	return TableNameSysDict
}
