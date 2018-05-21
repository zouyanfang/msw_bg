package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}



//模板
func (this *BaseController)IsneedTemplate(){
	this.Layout = "layout.html"
}
