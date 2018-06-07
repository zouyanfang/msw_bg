package controllers

import (
	"msw_bg/models"
	"time"
	"msw_bg/service"
	"msw_bg/util"
)

type MenuController struct {
	BaseController
}

//到菜单页
func(this *MenuController)ToMenu(){
	resp := service.GetMenuList(1,util.PAGE_SIZE,"")
	this.Data["object"] = resp
	this.Data["url"] = "/menu/pagemenu"
	this.IsneedTemplate()
	this.TplName = "menu.html"
}

func (this *MenuController)PageMenu(){
	var resp models.PageResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	name := this.GetString("name")
	page,_ :=this.GetInt("page")
	resp = service.GetMenuList(page,util.PAGE_SIZE,name)
}

func (this *MenuController)Delete(){
	var resp models.BaseResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id,err :=this.GetInt("menuid")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	err = models.DeleteMenu(id)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
}

//创建菜单
func (this *UserController)CreateNewImg(){
	var resp models.BaseResp
	resp.Ret = 403
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请先登入"
	}
	s,f,err := this.GetFile("img")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	id := int(time.Now().Unix())
	path := util.SaveImg(s,f,id,"menu_img/")
	path = "http://127.0.0.1:8080/static/img/menu_img/"+path
	menuname := this.GetString("menuname")
	describe := this.GetString("describe")
	resp = service.CreateNewMenu(this.User.Id,menuname,path,describe)
}

