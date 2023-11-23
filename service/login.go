package service

import (
	"github.com/gin-gonic/gin"
	"go-admin/define"
	"go-admin/helper"
	"go-admin/models"
	"net/http"
)

func LoginPassword(c *gin.Context) {
	in := new(LoginPasswordRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	sysUser, err := models.GetUserByUsernamePassword(in.UserName, in.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}

	token, err := helper.GenerateToken(sysUser.ID, sysUser.Username, define.TokenExpire)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return

	}

	refreshToken, err := helper.GenerateToken(sysUser.ID, sysUser.Username, define.RefreshExpire)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return

	}

	data := &LoginPasswordReply{
		Token:        token,
		RefreshToken: refreshToken,
		UserInfo:     sysUser,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": data,
	})
}
