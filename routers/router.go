package routers

import (
	"bigger/controllers"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/about", &controllers.BaseController{}, "*:About")
	beego.Router("/login", &controllers.LRController{}, "get:Login;post:LoginPost")
	beego.Router("/register", &controllers.LRController{}, "get:Register;post:RegisterPost")
	beego.Router("/my", &controllers.UserController{}, "get:Main")
	beego.Router("/my/photo", &controllers.PicController{}, "get:Add;post:SavePhoto")
	beego.Router("/api/my/photos", &controllers.PicController{}, "get:MyPhotos")
	beego.Router("/api/photos", &controllers.PhotoController{}, "get:Photo")
	beego.Router("/files", &controllers.FileController{}, "post:UploadFileByFrom")
	beego.Handler("/captcha/*", captcha.Server(captcha.StdWidth, captcha.StdHeight))
}
