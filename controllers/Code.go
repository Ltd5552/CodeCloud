package controllers

import (
	"codecloud/models/service"
	"codecloud/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

var codeService service.CodeService

// CodeRunController 代码运行控制器
type CodeRunController struct {
	beego.Controller
}

func (c *CodeRunController) Get() {
	c.TplName = "runner.html"
}

// Post  /code/run
func (c *CodeRunController) Post() {
	//获取参数
	code := c.GetString("code")
	Type := c.GetString("type")
	uid := c.GetString("userID")
	runTime := c.GetString("runTime")
	cid := utils.RandStr(7)

	//定义结果变量
	var result string

	//进行核心运行
	//运行go的docker
	if Type == "golang" {
		codeService.Save(cid, uid, code, result, runTime, Type)
		//运行python的docker
	} else if Type == "python" {
		codeService.Save(cid, uid, code, result, runTime, Type)
	} else {
		result = "请求错误"
	}

	//响应参数传递
	type response struct {
		Result string
	}
	var res response
	res.Result = result
	c.Data["json"] = res
	err := c.ServeJSON()
	if err != nil {
		fmt.Println(err)
	}

}

// CodeAllHistoryController CodeAllHistory 所有历史记录查看
type CodeAllHistoryController struct {
	beego.Controller
}

// Get /code/history
func (c *CodeAllHistoryController) Get() {
	//获取参数
	uid := c.GetString("userID")

	lists := codeService.GetList(uid)

	c.Data["json"] = lists

	err := c.ServeJSON()
	if err != nil {
		fmt.Println(err)
	}
}

// CodeDetailController 一次记录的详情
type CodeDetailController struct {
	beego.Controller
}

// Get /code/detail
func (c *CodeDetailController) Get() {
	cid := c.GetString("codeID")
	detail := codeService.GetDetail(cid)

	type response struct {
		Result  string
		Code    string
		Type    string
		RunTime string
	}
	var res = response{detail.Result, detail.Code, detail.Type, detail.Time}
	c.Data["json"] = res
	err := c.ServeJSON()
	if err != nil {
		fmt.Println(err)
	}
}

// Delete /code/detail
func (c *CodeDetailController) Delete() {
	cid := c.GetString("codeID")
	codeService.DeleteOne(cid)
}

// CodeRunAgainController CodeRunAgain 代码再次运行
type CodeRunAgainController struct {
	beego.Controller
}

// Get /code/runAgain
func (c *CodeRunAgainController) Get() {
	cid := c.GetString("codeID")
	type response struct {
		Code string
	}
	code := codeService.GetCode(cid)
	var res = response{code}
	c.Data["json"] = res
	err := c.ServeJSON()
	if err != nil {
		fmt.Println(err)
	}
}
