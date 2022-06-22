package dao

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

// 定义一个全局对象db
var db *sql.DB

func init() {

	user := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	database := beego.AppConfig.String("database")
	//Sprintf返回拼接结果
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", user, password, host, port, database)

	// 这里只是赋值，不会校验账号密码是否正确
	// 这里不要使用:=由于是给全局变量db赋值，然后在main函数中使用全局变量db
	db, _ = sql.Open("mysql", datasource)
	// 尝试与数据库建立连接（这里会校验dsn内容是否正确）
	err := db.Ping()
	if err != nil {
		return
	}

}
