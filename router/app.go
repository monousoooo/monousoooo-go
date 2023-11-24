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
	r.DELETE("/user", service.DeleteUser)
	return r
}
