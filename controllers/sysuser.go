package controllers

type SysUserController struct {
	BaseController
}

func (this *SysUserController)ToSys(){
	this.IsneedTemplate()
	this.TplName = "systemuser.html"
}

func (this *SysUserController)ToNews(){
	this.IsneedTemplate()
	this.TplName = "news.html"
}

