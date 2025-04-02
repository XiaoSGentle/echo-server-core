// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysDictChild = "sys_dict_child"

// SysDictChild mapped from table <sys_dict_child>
type SysDictChild struct {
	ID         int64          `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	DictCode   string         `gorm:"column:dict_code;type:varchar(255)" json:"dictCode"`
	Type       int64          `gorm:"column:type;type:int(11)" json:"type"`
	Value      string         `gorm:"column:value;type:varchar(255)" json:"value"`
	Label      string         `gorm:"column:label;type:varchar(255)" json:"label"`
	Style      string         `gorm:"column:style;type:varchar(255)" json:"style"`
	Describe   string         `gorm:"column:describe;type:varchar(255)" json:"describe"`
	OrderNum   int64          `gorm:"column:order_num;type:int(11)" json:"orderNum"`
	ItemClass  string         `gorm:"column:item_class;type:varchar(255)" json:"itemClass"`
	CreateDept int64          `gorm:"column:create_dept;type:int(11)" json:"createDept"`
	CreateBy   int64          `gorm:"column:create_by;type:int(11)" json:"createBy"`
	CreateTime time.Time      `gorm:"column:create_time;autoCreateTime;type:datetime" json:"createTime"`
	UpdateBy   int64          `gorm:"column:update_by;type:int(11)" json:"updateBy"`
	UpdateTime time.Time      `gorm:"column:update_time;autoUpdateTime;type:datetime" json:"updateTime"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:varchar(255)" json:"deleteTime"`
}

// TableName SysDictChild's table name
func (*SysDictChild) TableName() string {
	return TableNameSysDictChild
}
