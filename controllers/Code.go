package controllers

import (
	"codecloud/models/domain"
	"codecloud/models/service"
	beego "github.com/beego/beego/v2/server/web"
)

var code domain.Code
var codeService service.CodeService

// CodeRunController 代码运行控制器
type CodeRunController struct {
	beego.Controller
}

// Post  /code/run请求
func (c *CodeRunController) Post() {

}

// CodeAllHistoryController CodeAllHistory 所有历史记录查看
type CodeAllHistoryController struct {
	beego.Controller
}

// Get /code/allHistoryController
func (c *CodeAllHistoryController) Get() {

}

// CodeDetailController 一次记录的详情
type CodeDetailController struct {
	beego.Controller
}

// Get /code/oneDetail
func (c *CodeDetailController) Get() {

}

// Delete /code/oneDetail
func (c *CodeDetailController) Delete() {

}

// CodeRunAgainController CodeRunAgain 代码再次运行
type CodeRunAgainController struct {
	beego.Controller
}

// Get /code/runAgain
func (c *CodeRunAgainController) Get() {

}
