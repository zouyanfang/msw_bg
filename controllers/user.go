package controllers

import "msw_bg/service"

type UserController struct {
	BaseController
}

func (this *UserController)ToUser(){
	this.IsneedTemplate()
	nums := service.GetUserRegisterNum()
	this.Data["nums"] = nums
	this.TplName = "user.html"
	return
}

