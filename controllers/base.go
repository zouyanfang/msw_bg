package controllers

import (
	"github.com/astaxie/beego"
	"msw_bg/models"
	"fmt"
)

type BaseController struct {
	beego.Controller
	User *models.Admin
}

func (this *BaseController)Prepare(){
	name := this.Ctx.GetCookie("name")
	pwd :=this.Ctx.GetCookie("pwd")
	fmt.Println("name:",name,"pwd:",pwd)
	if name != "" && pwd != ""{
		user := models.Login(name,pwd)
		if user==nil{
			this.Redirect("/login",302)
			return
		}
		this.User = user
		this.Data["user"] = user
	}else{
		this.Redirect("/login",302)
	}
}

//模板
func (this *BaseController)IsneedTemplate(){
	this.Layout = "layout.html"
}
