package models

import (
	"bigger/utils"
	"errors"
	"fmt"
	"time"
)

type User struct {
	Id          int64  `xorm:"pk autoincr" json:"id"`
	UserName    string `orm:"unique" json:"username"`
	PassWord    string
	Email       string `orm:"unique" json:"email"`
	Pic         string `json:"pic"`
	VerifyEmail int    `json:"verifyEmail"`
	Phone       string `json:"phone"`
	Lever       int64  `json:"lever"`
	FirstIp     string
	LastIp      string
	Sina        string
	Blog        string
	Token       string    `json:"token"`
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
}

func (this *User) Add() error {
	this.PassWord = utils.EncodePasswd(PassWdSrc, this.PassWord)
	return insert(this)
}

func (this *User) Update() error {
	_, err := engine.Where("id=?", this.Id).Update(this)
	return err
}

func GetUserByToken(token string) (user *User, err error) {
	user = new(User)
	has, err := engine.Where("token=?", token).Get(user)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	if !has {
		return nil, errors.New("not found user.")
	}
	user.PassWord = ""
	return
}

func Login(userName, email, pwd string) (user *User, err error) {
	pwd = utils.EncodePasswd(PassWdSrc, pwd)
	user = new(User)
	has, err := engine.Where("(user_name=? OR email=?) AND pass_word=?", userName, email, pwd).Get(user)
	if err != nil {
		return
	}
	if !has {
		return nil, errors.New("not found user.")
	}
	user.Token = utils.EncodePasswd(PassWdSrc, fmt.Sprintf("%s_%d", userName, time.Now().Unix()))
	err = user.Update()
	if err != nil {
		return
	}
	user.PassWord = ""
	return
}
