package user

import (
	db "../../database"
	. "../../static"
	"fmt"
	"github.com/go-xorm/core"
	"reflect"
)

type User struct {
	Id      string `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Pwd     string `json:"pwd" form:"pwd"`
	Sex     int    `json:"sex" form:"sex"`
	OldYear int    `json:"oldYear" form:"old_year"`
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

func JudgeLogin(username string, pwd string) (b bool, i int,u *User) {
	var user User
	has, err := db.Orm.Where("name = ? and pwd = ?", username, pwd).Get(&user)
	fmt.Println(err)
	if err != nil {
		return false, 0,nil
	}
	if has == false {
		return false, 1,nil
	}
	return true, 1,&user
}

func UserRegister(param User)(b bool,str string) {
	var user User
	user.Id = MakeMD5(param.Name)
	user.Name=param.Name
	user.Pwd=param.Pwd
	user.OldYear=param.OldYear
	user.Sex=param.Sex
	user.Birth=param.Birth
	has,err := db.Orm.Insert(&user)
	if has == 0 {
		return false, "该用户已经注册!请不要重复操作"
	}
	if err != nil {
		return false,"系统错误!请稍后再试"
	}
	return true,"注册成功请登陆!"
}
