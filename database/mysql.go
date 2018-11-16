package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var Orm *xorm.Engine

func init() {
	var err error
	Orm, err = xorm.NewEngine("mysql", "root:zsjin0520@tcp(148.70.100.22:3306)/test?charset=utf8")
	Orm.ShowSQL(true)
	if err != nil {
		log.Fatal("err:", err.Error())
	} else {
		//log.Fatal("then:","数据库连接成功!")
	}

}
