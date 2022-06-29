package controllers

import (
	"codecloud/domain"
	"codecloud/models/service"
	"codecloud/utils"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

var codeService service.CodeService
var apiService service.ApiService

// CodeRunController 代码运行控制器
type CodeRunController struct {
	beego.Controller
}

// Get  /code/run
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
	//随机cid
	cid := utils.RandStr(7)

	var res domain.Res
	//运行go的docker
	if Type == "Golang" {
		res = apiService.Golang(code, res)
		codeService.Save(cid, uid, code, res.Result, res.Errors, runTime, Type)
		//运行python的docker
	} else if Type == "Python" {
		res = apiService.Python(code, res)
		codeService.Save(cid, uid, code, res.Result, res.Errors, runTime, Type)
	}
	//响应参数传递
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
	type response struct {
		Lists []domain.List
	}
	uid := c.GetString("userID")
	var res response
	res.Lists = codeService.GetList(uid)
	//for _, i := range res.Lists {
	//	fmt.Println(i)
	//}
	//fmt.Println(res.lists)
	c.Data["json"] = res

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
	c.TplName = "history.html"
}
func (c *CodeDetailController) Post() {
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
	fmt.Println("cid =", cid)
	type response struct {
		Flag bool
	}
	var res response
	if cid != "" {
		res.Flag = codeService.DeleteOne(cid)
	} else {
		res.Flag = false
	}
	c.Data["json"] = res
	err := c.ServeJSON()
	if err != nil {
		return
	}

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
