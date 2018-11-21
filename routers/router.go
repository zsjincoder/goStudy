package routers

import (
	. "../controller"
	"fmt"
	"github.com/fwhezfwhez/jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(Cors())
	router.POST("login", UserLogin)
	router.POST("register", UserReg)
	router.Use(Validate())
	router.GET("getUser", GetUser)
	router.GET("getAllUsers", GetUsers)
	router.GET("queryDbInfo", QueryDbInfo)
	router.GET("querySomeCols", GetSomeCols)
	return router
}

/**
  *2018/11/21
  *author:xiaoC
  *func:处理跨域问题
  *param:
 */

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

/**
  *2018/11/21
  *author:xiaoC
  *func:处理用户权限
  *param:
 */
func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if JWTToken := c.Request.Header.Get("Authorization"); JWTToken != "" {
			token := jwt.GetToken()
			legal, err := token.IsLegal(JWTToken, "1932019050")
			if err != nil {
				fmt.Println(err)
				c.Abort()
				c.JSON(200, gin.H{
					"data":"响应令牌验证错误",
				})
				return
			}
			if !legal {
				c.Abort()
				c.JSON(200, gin.H{
					"data":"响应令牌验证错误",
				})
				return
			}
			c.Next()
		} else {
			c.JSON(200, gin.H{
				"data":"未找到令牌",
			})
			c.Abort()
			return
		}
	}
}
