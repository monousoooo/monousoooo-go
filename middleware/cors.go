package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")  // 允许访问所有域
		c.Header("Access-Control-Allow-Methods", "*") // 允许访问所有域
		c.Header("Access-Control-Allow-Headers", "*") // 允许访问所有域
		c.Next()
	}
}
