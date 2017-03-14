package main

import (
	"bigger/models"
	_ "bigger/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.SyncModels()
	beego.SetStaticPath("/upload", "upload")
	beego.Run()
}
