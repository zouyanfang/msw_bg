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
	resp := service.GetUser(1,util.PAGE_SIZE)
	fmt.Println(resp)
	this.Data["object"] = resp
	this.Data["url"] = "/user/pageuser"
	this.IsneedTemplate()
	nums := service.GetUserRegisterNum()
	this.Data["nums"] = nums
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