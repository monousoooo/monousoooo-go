package models

import "gorm.io/gorm"

type SysRole struct {
	gorm.Model
	Name    string `gorm:"column:name;type:varchar(100);" json:"name"`
	IsAdmin int8   `gorm:"column:is_admin;type:tinyint(1);default:0;" json:"is_admin"`
	Sort    int64  `gorm:"column:sort;type:int(11);default:0;" json:"sort"`
	Remarks string `gorm:"column:remarks;type:varchar(255);" json:"remarks"`
}

func (table *SysRole) TableName() string {
	return "sys_role"
}
