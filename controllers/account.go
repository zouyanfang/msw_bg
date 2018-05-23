package controllers

import (
	"github.com/astaxie/beego"
	"msw_bg/models"
)

type AccountController struct {
	beego.Controller
}

//到登入页面上去
func (this *AccountController)ToLogin(){
	this.TplName = "login.html"
}

//登录
func (this *AccountController)Login(){
	var resp models.BaseResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	name := this.GetString("name")
	pwd := this.GetString("pwd")
	u:= models.Login(name,pwd)
	if u == nil {
		return
	}
	this.Ctx.SetCookie("name",u.Account)
	this.Ctx.SetCookie("pwd",u.Pwd)
	resp.Ret = 200
}

//登出
func (this *AccountController)LoginOut(){
	this.Ctx.SetCookie("name","")
	this.Ctx.SetCookie("pwd","")
	this.Redirect("/login",302)
}


