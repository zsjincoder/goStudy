package controller

import (
	. "../models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

/**
  *2018/11/15
  *author:xiaoC
  *func:查询一条信息
  *param:
 */
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
func GetUsers(c *gin.Context)  {
	user:= GetAllUser()
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
func QueryDbInfo(c *gin.Context)  {
	info := GetDbMates()
	if info !=nil{
		c.JSON(http.StatusOK,gin.H{
			"info":info,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"info":nil,
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
		c.JSON(http.StatusOK,gin.H{
			"msg":"数据库查询内如为空!",
		})
	}else {
		c.JSONP(http.StatusOK,gin.H{
			"msg":"操作成功",
			"data":info,
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
	UserName string
	Password string
}
 func UserLogin (c *gin.Context){
 	var userInfo UserInfo
 	err := c.ShouldBindJSON(&userInfo)
 	fmt.Println(&userInfo)
	 if err != nil {
		 c.JSON(http.StatusOK,gin.H{
		 	"err":"服务器内部错误!",
		 })
	 }else {
	 	user :=userInfo.UserName
	 	pwd := userInfo.Password
	 	isExist := JudgeLogin(user,pwd)
		 if isExist{
			 c.JSON(http.StatusOK,gin.H{
				 "msg":"操作成功!",
				 "data":"请尽情享受吧!",
			 })
		 }else {
			 c.JSON(http.StatusOK,gin.H{
				 "msg":"操作失败!",
				 "data":"用户名密码错误!请检查!",
			 })
		 }
	 }

 }