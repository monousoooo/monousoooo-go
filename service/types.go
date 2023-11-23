package service

import "go-admin/models"

type LoginPasswordRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginPasswordReply struct {
	Token        string          `json:"token"`
	RefreshToken string          `json:"refresh_token"`
	UserInfo     *models.SysUser `json:"userInfo"`
}
