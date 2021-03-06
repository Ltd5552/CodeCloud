package routers

import (
	"codecloud/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserLoginController{})
	beego.Router("/user/findName", &controllers.UserFindNameController{})
	beego.Router("/code/run", &controllers.CodeRunController{})
	beego.Router("/code/history", &controllers.CodeAllHistoryController{})
	beego.Router("/code/detail", &controllers.CodeDetailController{})
	beego.Router("/code/runAgain", &controllers.CodeRunAgainController{})
}
