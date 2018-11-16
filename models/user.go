package models

import (
	db "../database"
	"fmt"
	"github.com/go-xorm/core"
	"reflect"
)

type User struct {
	Id      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Pwd     string `json:"pwd" form:"pwd"`
	Sex     int    `json:"sex" form:"sex"`
	OldYear int    `json:"old_year" form:"old_year"`
	Birth   string `json:"birth" form:"birth"`
}

/**
  *2018/11/15
  *author:xiaoC
  *func:查询user表中第一条数据
  *param:
 */
func GetUserByOne() (u *User) {
	var user User
	has, err := db.Orm.Get(&user)
	fmt.Println(err)
	if err != nil {
		//log.Fatal(err.Error())
		return nil
	}
	if has == false {
		return nil
	}
	return &user
}

func GetAllUser() (u []User) {
	user := make([] User, 0)
	//user:=make(map[int64] *User)
	err := db.Orm.Find(&user)
	if err != nil {
		return nil
		panic(err)
	} else {
		return user
	}
}

func GetDbMates() []*core.Table {
	info, err := db.Orm.DBMetas()
	fmt.Println(reflect.TypeOf(info))
	if err != nil {
		return nil
	} else {
		return info
	}
}

func GetSomeColsInUser() (u []User) {
	user := make([]User, 0)
	err := db.Orm.Cols("name", "sex").Find(&user)
	fmt.Println(reflect.TypeOf(err))
	if err != nil {
		return nil
		panic(err)
	} else {
		return user
	}
}

func JudgeLogin(username string, pwd string) (b bool) {
	var user User
	has,err := db.Orm.Where("name = ? and pwd = ?", username, pwd).Get(&user)
	fmt.Println(err)
	if err != nil {
		return false
	}
	if has==false{
		return  false
	}
	return true
}