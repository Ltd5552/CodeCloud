package controllers

import (
	"codecloud/models/domain"
	"codecloud/models/service"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

var UserService service.UserService
var user domain.User

// UserLoginController 登录注册请求
type UserLoginController struct {
	beego.Controller
}

// Get /user/RegisterOrLogin
func (u *UserLoginController) Get() {
	u.TplName = "valid.html"
}

// Post /user/RegisterOrLogin
func (u *UserLoginController) Post() {
	u.TplName = "valid.html"
	name := u.GetString("name")
	password := u.GetString("password")
	fmt.Println(name)
	fmt.Println(password)

}

// UserFindNameController 获取用户名
type UserFindNameController struct {
	beego.Controller
}

// Get /user/findName
func (u *UserFindNameController) Get() {
	name := u.GetString("name")
	UserService.FindName(name)
}
