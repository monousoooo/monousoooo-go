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

func GetRoleList(keyword string) *gorm.DB {
	tx := DB.Model(new(SysRole)).Select("id,name,is_admin,sort,remarks,created_at,updated_at")
	if keyword != "" {
		tx.Where("name LIKE ?", "%"+keyword+"%")
	}
	tx.Order("sort ASC")
	return tx
}

func GetRoleDetail(id uint) (*SysRole, error) {
	data := new(SysRole)
	err := DB.Model(new(SysRole)).Where("id = ?", id).First(data).Error
	return data, err
}
