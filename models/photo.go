package models

import (
	"time"
)

type Photo struct {
	Id        int64     `xorm:"pk autoincr" json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Pic       string    `json:"pic"`
	UserId    int64     `json:"userId"`
	Love      int       `json:"love"` //喜欢人数
	View      int       `json:"view"` //浏览人数
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (this *Photo) Add() error {
	return insert(this)
}

func PhotoList() (r []*Photo, err error) {
	err = engine.Find(&r)
	return
}

func MyPhotoList(uid int64) (r []*Photo, err error) {
	err = engine.Where("user_id=?", uid).Find(&r)
	return
}
