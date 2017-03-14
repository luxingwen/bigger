package controllers

import (
	"bigger/models"
	"fmt"
)

type AppController struct {
	BaseController
	User *models.User
}

func (this *AppController) Prepare() {

	user := this.GetSession("user")
	if user == nil {
		token := this.GetString("token", "")
		if token != "" {
			user, err := models.GetUserByToken(token)
			if err != nil {
				fmt.Println("getuser bytoken err. ", err)
				this.Data["json"] = map[string]interface{}{"code": 1, "msg": "无效用户，请登陆"}
				this.ServeJSON()
				this.StopRun()
			} else {
				this.SetSession("user", user)
			}
		}
	}
	if v, ok := user.(*models.User); ok {
		this.User = v
	}
	if this.User == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "msg": "无效用户，请登陆"}
		this.ServeJSON()
		this.StopRun()
	}
	this.Data["IsLogin"] = true
	this.Data["User"] = this.User
	fmt.Println("user--->", this.User)
}
