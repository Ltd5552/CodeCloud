package controllers

import (
	"codecloud/models/service"
	"codecloud/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

var UserService service.UserService

// UserLoginController 登录注册请求
type UserLoginController struct {
	beego.Controller
}

// Get /user/login
func (u *UserLoginController) Get() {
	//请求路径 /user/login获取的页面
	u.TplName = "regist_login.html"
}

// Post /user/login
func (u *UserLoginController) Post() {
	//下一步请求的地址
	//u.TplName = "runner.html"

	//获取请求参数
	name := u.GetString("username")
	password := u.GetString("password")
	Type := u.GetString("type")

	fmt.Println(name, password, Type)

	type response struct {
		Flag   int
		Msg    string
		UserID string
	}

	var res response
	//登录
	if Type == "1" {
		userID, flag := UserService.CheckAccount(name, password)
		if flag {
			//	登录成功，返回userID、flag为1，msg置空，
			res.UserID = userID
			res.Flag = 1
			res.Msg = ""
		} else {
			//	登录失败
			res.Flag = 0
			res.UserID = ""
			res.Msg = "登录失败"
			fmt.Println("登录失败")
		}
		//注册
	} else if Type == "0" {
		uid := utils.RandStr(7)
		fmt.Println(uid)
		err := UserService.Save(uid, name, password)
		if err != nil {
			//	注册失败，用户名已经存在
			fmt.Println("注册失败")
			res.Flag = 0
			res.Msg = "注册失败"
			res.UserID = ""
		}
		res.Flag = 1
		res.Msg = "注册成功"
		res.UserID = uid
	}
	//返回数据,json格式
	u.Data["json"] = res
	err := u.ServeJSON()
	if err != nil {
		fmt.Println(err)
	}

}

// UserFindNameController 获取用户名
type UserFindNameController struct {
	beego.Controller
}

// Get /user/findName
func (u *UserFindNameController) Get() {
	uid := u.GetString("userID")
	name := UserService.FindName(uid)
	type response struct {
		UserName string
	}
	var res response
	if name != "" {
		res.UserName = name
	}
	//返回数据，json格式
	u.Data["json"] = res
	err := u.ServeJSON()
	if err != nil {
		fmt.Println(err)
	}
}
