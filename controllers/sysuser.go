package controllers

type SysUserController struct {
	BaseController
}

func (this *SysUserController)ToSys(){
	if this.User.Role != 1 {
		this.Ctx.WriteString("你没有该权限")
		return
	}
	this.IsneedTemplate()
	this.TplName = "systemuser.html"
}

func (this *SysUserController)ToNews(){
	this.IsneedTemplate()
	this.TplName = "news.html"
}


//登出
func (this *SysUserController)LoginOut(){
	this.Ctx.SetCookie("name","")
	this.Ctx.SetCookie("pwd","")
	this.Redirect("/login",302)
}