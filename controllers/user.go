package controllers


import (
	"msw_bg/service"
	"msw_bg/util"
	"fmt"
	"msw_bg/models"
)

type UserController struct {
	BaseController
}

func (this *UserController)ToUser(){
	this.IsneedTemplate()
	resp := service.GetUser(1,util.PAGE_SIZE)
	fmt.Println(resp)
	this.Data["object"] = resp
	this.Data["url"] = "/user/pageuser"
	//注册人统计
	nums := service.GetUserRegisterNum()
	//留言人数统计
	counts := service.GetUserMessageNum()
	this.Data["nums"] = nums
	this.Data["counts"] = counts
	this.TplName = "user.html"
	return
}

//用户留言分页
func (this *UserController)PageUser(){
	var resp models.PageResp
	resp.Ret = 403
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	page,err :=this.GetInt("page")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp = service.GetUser(page,util.PAGE_SIZE)
}


func (this *UserController)CheckDetail(){
	var resp models.BaseResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id,_ := this.GetInt("id")
	m,err :=models.GetUserMsgDetail(id)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Object = m
	resp.Ret = 200
}