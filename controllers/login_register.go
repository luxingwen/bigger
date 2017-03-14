package controllers

import (
	"bigger/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	//"net/http"
	"math/rand"
)

type LRController struct {
	BaseController
}

func (this *LRController) Login() {
	this.Data["CaptchaId"] = captcha.New()
	this.TplName = "login.html"
}

func (this *LRController) Register() {
	this.Data["CaptchaId"] = captcha.New()
	this.TplName = "register.html"
}

func (this *LRController) LoginPost() {
	var reqUser struct {
		Email           string `json:"email"`
		PassWord        string `json:"password"`
		CaptchaSolution string `json:"captchaSolution"`
		CaptchaId       string `json:"captchaId"`
	}
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	err := json.Unmarshal(body, &reqUser)
	if err != nil {
		this.Fail(400, err.Error())
	}

	if !verifyCode(reqUser.CaptchaId, reqUser.CaptchaSolution) {
		this.Fail(400, "验证码不正确。")
	}

	user, err := login(reqUser.Email, reqUser.Email, reqUser.PassWord)
	if err != nil {
		this.Fail(400, "用户名或密码不正确。")
	}

	this.SetSession("User", user)
	this.Succuess(user)

}
func login(username, email, password string) (*models.User, error) {
	u, err := models.Login(username, email, password)
	if err != nil {
		return nil, err
	}
	return u, nil
}

var Pic []string = []string{"http://a4.topitme.com/l/201010/16/12872170172866.jpg", "http://a4.topitme.com/l/201008/24/12826374341928.jpg", "https://avatars3.githubusercontent.com/u/16460296?v=3&s=460"}

func (this *LRController) RegisterPost() {
	var reqUser struct {
		UserName        string `json:"username"`
		Email           string `json:"email"`
		PassWord        string `json:"password"`
		CaptchaSolution string `json:"captchaSolution"`
		CaptchaId       string `json:"captchaId"`
	}

	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	err := json.Unmarshal(body, &reqUser)
	if err != nil {
		this.Fail(400, err.Error())
	}

	if !verifyCode(reqUser.CaptchaId, reqUser.CaptchaSolution) {
		this.Fail(400, "验证码不正确。")
	}

	user := &models.User{UserName: reqUser.UserName, PassWord: reqUser.PassWord, Email: reqUser.Email}
	user.Pic = Pic[rand.Intn(len(Pic))]
	err = user.Add()
	if err != nil {
		fmt.Println(err)
		this.Fail(400, err.Error())
	}
	user.PassWord = ""
	this.SetSession("user", user)
	this.Succuess(user)
}

func verifyCode(captchaId, captchaSolution string) bool {
	return captcha.VerifyString(captchaId, captchaSolution)
}
