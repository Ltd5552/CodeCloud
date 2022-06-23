package dao

import (
	"database/sql"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

// 定义一个全局对象db
var db *sql.DB

// 初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	user, _ := beego.AppConfig.String("username")
	password, _ := beego.AppConfig.String("password")
	host, _ := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.String("port")
	database, _ := beego.AppConfig.String("database")
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", user, password, host, port, database)
	// 这里只是赋值，不会校验账号密码是否正确
	// 这里不要使用:=由于是给全局变量db赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", datasource)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（这里会校验dsn内容是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
