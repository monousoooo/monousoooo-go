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

type GetUserListRequest struct {
	*QueryRequest
}

type QueryRequest struct {
	Page    int    `json:"page" form:"page"`
	Size    int    `json:"size" form:"size"`
	Keyword string `json:"keyword" form:"keyword"`
}

type GetUserListReply struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Remarks  string `json:"remarks"`
}

type UpdateUserRequest struct {
	ID uint `json:"id"`
	AddUserRequest
}
