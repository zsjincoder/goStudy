package controller

import (
	. "../models/user"
	"fmt"
	"github.com/fwhezfwhez/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

/**
  @time:2018/11/15
  @author:xiaoC
  @func:查询一条信息
  @param:
 */

var secretKey = "1932019050"

func GetUser(c *gin.Context) {
	user := GetUserByOne()
	if user == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true, "data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true, "data": user,
		})
	}
}

/**
  *2018/11/15
  *author:xiaoC
  *func:查询user表全部信息
  *param:
 */
func GetUsers(c *gin.Context) {
	user := GetAllUser()
	if user == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true, "data": "",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true, "data": user,
		})
	}
}

/**
  *2018/11/15
  *author:xiaoC
  *func:查询数据库信息
  *param:
 */
func QueryDbInfo(c *gin.Context) {
	info := GetDbMates()
	if info != nil {
		c.JSON(http.StatusOK, gin.H{
			"info": info,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"info": nil,
		})
	}
}

/**
  *2018/11/15
  *author:xiaoC
  *func:根据name和sex查询user
  *param:
 */
func GetSomeCols(c *gin.Context) {
	info := GetSomeColsInUser
	fmt.Println(reflect.TypeOf(info))
	if info == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "数据库查询内如为空!",
		})
	} else {
		c.JSONP(http.StatusOK, gin.H{
			"msg":  "操作成功",
			"data": info,
		})
	}
}

/**
  *2018/11/15
  *author:xiaoC
  *func:登陆接口
  *param:
 */
type UserInfo struct {
	Id      string `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Pwd     string `json:"pwd" form:"pwd"`
	Sex     int    `json:"sex" form:"sex"`
	OldYear int    `json:"oldYear" form:"oldYear"`
	Birth   string `json:"birth" form:"birth"`
}

func UserLogin(c *gin.Context) {
	//var userInfo UserInfo
	//err := c.ShouldBindJSON(&userInfo)
	name := c.PostForm("userName")
	pwd := c.PostForm("password")
	if name == "" || pwd == "" {
		c.JSON(http.StatusOK, gin.H{
			"err": "服务器内部错误!",
		})
	} else {
		//name := userInfo.UserName
		//pwd := userInfo.Password
		isExist, needReg,u := JudgeLogin(name, pwd)
		if isExist && needReg == 1 {
			//获取token管理对象
			token := jwt.GetToken()
			//添加令牌关键信息
			token.AddPayLoad("userName", name).AddPayLoad("role", "admin").AddHeader("typ", "JWT").AddHeader("alg", "HS256")
			//添加令牌期限
			exp := time.Now().Add(1 * time.Hour)
			token.AddPayLoad("exp", strconv.FormatInt(exp.Unix(), 10))
			//获取令牌，并添加进reponse的header里
			jwts, _, erre := token.JwtGenerator(secretKey)
			if erre != nil {
				fmt.Println("token生成出错")
				return
			}
			c.Writer.Header().Add("x-auth-token", jwts)
			c.JSON(http.StatusOK, gin.H{
				"Tips":      "操作成功!",
				"code":200,
				"msg":       "请尽情享受吧!",
				"data":      "toLogin",
				"areYouReg": true,
				"token":     jwts,
				"user":u,
			})
		} else if isExist && needReg == 0 {
			c.JSON(http.StatusOK, gin.H{
				"Tips":      "操作成功!",
				"code":200,
				"msg":       "用户名不存在，请前去注册",
				"data":      "toRegister",
				"areYouReg": false,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Tips": "操作失败!",
				"code":501,
				"msg":  "服务器内部错误!",
			})
		}
	}

}

/**
  *2018/11/16
  *author:xiaoC
  *func:用户注册
  *param:
 */
func UserReg(c *gin.Context) {
	var userInfo User
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		fmt.Println(&userInfo)
	} else {
		isSC, s := UserRegister(userInfo)
		fmt.Println(&userInfo)
		fmt.Println(isSC)
		if isSC {
			c.JSON(http.StatusOK, gin.H{
				"Tips":    "操作成功!",
				"code":200,
				"success": true,
				"msg":     s,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Tips":    "操作失败!",
				"code":200,
				"success": false,
				"msg":     s,
			})
		}
	}
}
