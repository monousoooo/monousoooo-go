package models

import "gorm.io/gorm"

type SysMenu struct {
	gorm.Model
	ParentId      uint   `gorm:"column:parent_id;type:int(11);" json:"parent_id"`
	Name          string `gorm:"column:name;type:varchar(100);" json:"name"`
	Path          string `gorm:"column:path;type:varchar(100);" json:"path"`
	Icon          string `gorm:"column:icon;type:varchar(100);" json:"icon"`
	Sort          int64  `gorm:"column:sort;type:int(11);default:0;" json:"sort"`
	ComponentName string `gorm:"column:component_name;type:varchar(100);" json:"component_name"`
}

func (table *SysMenu) TableName() string {
	return "sys_menu"
}

func GetMenuList() *gorm.DB {
	tx := DB.Model(new(SysMenu)).Select("id,parent_id,name,path,icon,sort,component_name,created_at,updated_at").Order("sort ASC")
	return tx
}
