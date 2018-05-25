package controllers

import (
	"msw_bg/service"
	"msw_bg/util"
	"msw_bg/models"
	"fmt"
)

type NewsController struct {
	BaseController
}

func (this *NewsController)ToNews(){
	resp := service.GetNews(1,util.PAGE_SIZE)
	this.Data["object"] = resp
	fmt.Println(resp)
	this.Data["url"] = "/news/pagenews"
	this.IsneedTemplate()
	this.TplName = "news.html"
}

func (this *NewsController)PageNews(){
	var resp models.PageResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	page,_ := this.GetInt("page")
	resp = service.GetNews(page,util.PAGE_SIZE)
	return
}

func (this *NewsController)Delete(){
	var resp models.BaseResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id,err :=this.GetInt("id")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	err = models.DeleteNews(id)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
}

//发布新闻
func (this *NewsController)ReleaseNews(){
	var resp models.BaseResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	title := this.GetString("title")
	if title == ""{
		resp.Msg = "标题不能为空"
		return
	}
	content := this.GetString("content")
	if content == ""{
		resp.Msg = "内容不能空"
		return
	}
	resp = service.ReleaseNews(this.User.Id,title,content)
}


