package middleware

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/common/ctype"
	"gvb_server/models/common/response"
	"gvb_server/service/redis"
	"gvb_server/utils/jwt"
)

// JwtAuth 判断是否登录
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		// 判断是否在redis中
		if redis.CheckLogout(token) {
			response.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		// 登录的用户
		c.Set("claims", claims)
	}
}

// JwtAdmin 判断是否为管理员
func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		// 判断是否在redis中
		if redis.CheckLogout(token) {
			response.FailWithMessage("token已失效", c)
			c.Abort()
			return
		}
		// 登录的用户
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			response.FailWithMessage("权限错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
