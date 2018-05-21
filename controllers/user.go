package controllers

type UserController struct {
	BaseController
}

func (this *UserController)ToUser(){
	this.IsneedTemplate()
	this.TplName = "user.html"
	return
}
