package controllers

import (
	"bigger/models"
)

type UserController struct {
	AppController
}

func (this *UserController) NestPrepare() {
	if this.User == nil {
		user := this.GetSession("user")
		if user == nil {
			this.Data["json"] = map[string]interface{}{"code": 1, "msg": "无效用户，请登陆"}
			this.ServeJSON()
			this.StopRun()
		}
		if v, ok := user.(*models.User); ok {
			this.User = v
		}
	}
}

func (this *UserController) Main() {
	this.TplName = "user.html"
}

func (this *UserController) UserInfo() {

}
