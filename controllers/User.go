package controllers

import (
	"codecloud/models/service"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

var UserService service.UserService

// UserLoginController 登录注册请求
type UserLoginController struct {
	beego.Controller
}

// Get /user/RegisterOrLogin
func (u *UserLoginController) Get() {
	//请求路径 /user/login获取的页面
	u.TplName = "valid.html"
}

// Post /user/RegisterOrLogin
func (u *UserLoginController) Post() {
	//下一步请求的地址
	u.TplName = "valid.html"

	//获取请求参数
	name := u.GetString("username")
	password := u.GetString("password")
	Type := u.GetString("type")

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
			err := u.SetSession("user", "")
			if err != nil {
				fmt.Println("Set session failed:", err)
			}
		} else {
			//	登录失败
			res.Flag = 0
			res.UserID = ""
			res.Msg = "登录失败"
		}
		//注册
	} else if Type == "0" {
		err := UserService.Save(name, password)
		if err != nil {
			//	注册失败，用户名已经存在
			res.Flag = 0
			res.Msg = "注册失败"
			res.UserID = ""
		}
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
	if name == "null" {
		//	不存在该用户，返回admin
		res.UserName = "admin"
	} else {
		//	返回该用户名
		res.UserName = name
	}

	//返回数据，json格式
	u.Data["json"] = res
	err := u.ServeJSON()
	if err != nil {
		fmt.Println(err)
	}
}
