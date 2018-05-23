package controllers

type MenuController struct {
	BaseController
}

//到菜单页
func(this *MenuController)ToMenu(){
	this.IsneedTemplate()
	this.TplName = "menu.html"
}


