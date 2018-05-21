package controllers

import "github.com/astaxie/beego"

type AccountController struct {
	beego.Controller
}

//到登入页面上去
func (this *AccountController)ToLogin(){
	this.TplName = "login.html"
}
