package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/middleware"
	"go-admin/service"
)

func App() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.POST("/login/password", service.LoginPassword)
	r.GET("/user", service.GetUserList)
	r.POST("/user", service.AddUser)
	r.PUT("/user", service.UpdateUser)
	r.DELETE("/user/:id", service.DeleteUser)
	r.GET("/role", service.GetRoleList)
	r.POST("/role", service.AddRole)
	r.GET("/role/detail/:id", service.GetRoleDetail)
	r.PUT("/role", service.UpdateRole)
	r.DELETE("/role/:id", service.DeleteRole)
	return r
}
