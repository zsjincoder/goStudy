package routers

import (
	. "../controller"
	"fmt"
	"github.com/fwhezfwhez/jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.POST("login", UserLogin)
	router.Use(Validate())
	router.POST("register", UserReg)
	router.GET("getUser", GetUser)
	router.GET("getAllUsers", GetUsers)
	router.GET("queryDbInfo", QueryDbInfo)
	router.GET("querySomeCols", GetSomeCols)
	return router
}

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if JWTToken := c.Request.Header.Get("x-auth-token"); JWTToken != "" {
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
