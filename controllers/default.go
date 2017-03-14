package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	user := this.GetSession("user")
	if user == nil {
		this.Data["IsLogin"] = false
	} else {
		this.Data["IsLogin"] = true
	}
}

func (c *BaseController) Get() {
	c.TplName = "index.html"
}

func (this *BaseController) About() {
	this.TplName = "about.html"
}

func (this *BaseController) Fail(code int, msg string) {
	this.Data["json"] = map[string]interface{}{"code": code, "msg": msg}
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) Succuess(data interface{}) {
	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "success", "data": data}

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("marshal err. ", err)
	}
	fmt.Println("res: ", string(b))
	this.ServeJSON()

}
