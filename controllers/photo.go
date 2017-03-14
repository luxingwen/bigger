package controllers

import (
	"bigger/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type PhotoController struct {
	BaseController
}

type PicController struct {
	AppController
}

func (this *PhotoController) Photo() {
	photos, err := models.PhotoList()
	if err != nil {
		this.Fail(400, err.Error())
	}
	this.Succuess(photos)
}

func (this *PicController) MyPhotos() {
	photos, err := models.MyPhotoList(this.User.Id)
	if err != nil {
		this.Fail(400, err.Error())
	}
	this.Succuess(photos)
}

func (this *PicController) Add() {
	this.TplName = "upload.html"
}

func (this *PicController) SavePhoto() {
	var req struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
		Pic  string `json:"pic"`
	}
	body := this.Ctx.Input.CopyBody(beego.BConfig.MaxMemory)
	err := json.Unmarshal(body, &req)
	if err != nil {
		this.Fail(400, err.Error())
	}

	photo := &models.Photo{Name: req.Name, Desc: req.Desc, Pic: req.Pic, UserId: this.User.Id}
	err = photo.Add()
	if err != nil {
		this.Fail(400, err.Error())
	}
	this.Succuess(photo)
}
