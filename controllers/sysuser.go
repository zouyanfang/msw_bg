package controllers

import (
	"msw_bg/service"
	"msw_bg/util"
	"msw_bg/models"
	"fmt"
)

type SysUserController struct {
	BaseController
}

func (this *SysUserController)ToSys(){
	if this.User.Role != 1 {
		this.Ctx.WriteString("你没有该权限")
		return
	}
	resp := service.GetAdmin(1,util.PAGE_SIZE)
	this.Data["object"] = resp
	this.Data["url"] = "/sysuser/pageadmin"
	this.IsneedTemplate()
	this.TplName = "systemuser.html"
}

func (this *SysUserController)PageAdmin(){
	var resp models.PageResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	page,_ := this.GetInt("page")
	resp = service.GetAdmin(page,util.PAGE_SIZE)
}

func (this *SysUserController)AddAdmin(){
	var resp models.BaseResp
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User.Role != 1{
		this.Ctx.WriteString("你没有该权限")
	}
	fmt.Println(1)
	account := this.GetString("account")
	if account == ""{
		resp.Msg = "账户不能为空"
		return
	}
	pwd := this.GetString("pwd")
	if pwd == "" {
		resp.Msg = "密码不能为空"
		return
	}
	resp = service.AddAdmin(account,pwd)
}

func (this *SysUserController)Delete(){
	var resp models.BaseResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id,err :=this.GetInt("id")
	fmt.Println(id)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	err = models.DeleteAdmin(id)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
}

func (this *SysUserController) UpdateSysUserState()  {
	var resp models.BaseResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id ,_:=this.GetInt("id")
	state,_ :=this.GetInt("state")
	fmt.Println("state",state)
	resp = service.UpdateSysUser(id,state)
}

//登出
func (this *SysUserController)LoginOut(){
	this.Ctx.SetCookie("name","")
	this.Ctx.SetCookie("pwd","")
	this.Redirect("/login",302)
}