package models

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	Username  string `gorm:"column:username;type:varchar(20);" json:"username"`
	Password  string `gorm:"column:password;type:varchar(20);" json:"password"`
	Phone     string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	WxUnionId string `gorm:"column:wx_union_id;type:varchar(20);" json:"wxUnionId"`
	WxOpenId  string `gorm:"column:wx_open_id;type:varchar(20);" json:"wxOpenId"`
	Avatar    string `gorm:"column:avatar;type:varchar(20);" json:"avatar"`
}

// TableName 设置表名称
func (table *SysUser) TableName() string {
	return "sys_user"
}

func GetUserByUsernamePassword(username, password string) (*SysUser, error) {
	data := new(SysUser)
	err := DB.Where("username = ? AND password = ?", username, password).First(data).Error
	return data, err
}
