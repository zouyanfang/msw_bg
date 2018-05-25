package service

import (
	"msw_bg/models"
	"msw_bg/util"
	"fmt"
)

//获取用户列表
func GetUser(page,pageSize int)(resp models.PageResp){
	start := util.StartIndex(page,pageSize)
	list ,err := models.UserList(start,pageSize)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	count,err := models.UserCount()
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Page = page
	ParsePage(&resp,count,pageSize,list)
	return
}

//获取公告
func GetNews(page,pageSize int)(resp models.PageResp){
	start := util.StartIndex(page,pageSize)
	list,err := models.GetNewsList(start,pageSize)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	count,err := models.CountNews()
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	for i,_ := range list{
		rs := []rune(list[i].Content)
		if len(rs)>12 {
			list[i].Content = string(rs[:12])
		}
	}
	resp.Page = page
	ParsePage(&resp,count,pageSize,list)
	return
}

//发布公告
func ReleaseNews(uid int,title,content string)(resp models.BaseResp){
	err := models.ReleaseNews(uid,title,content)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
	return
}

func GetAdmin(page,pageSize int)(resp models.PageResp){
	start := util.StartIndex(page,pageSize)
	list,err := models.GetAdminList(start,pageSize)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	count,err := models.CountAdmin()
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Page = page
	ParsePage(&resp,count,pageSize,list)
	return
}

func AddAdmin(account,pwd string)(resp models.BaseResp){
	count,err:= models.FindAccount(account)
	if err != nil{
		resp.Msg = err.Error()
		return
	}
	if count != 0 {
		resp.Msg = "用户已存在"
		return
	}
	err = models.InsertAdmin(account,pwd)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
	return
}


//可复用分页
func ParsePage(resp *models.PageResp,count,pageSize int,object interface{}){
	pages := util.PageCount(pageSize,count)
	if util.IsHaveNext(resp.Page,pages){

	}else if resp.Page >pages {
		resp.Page = 1
	}
	if resp.Page >= pages{
		resp.Next = false
	}else {
		resp.Next = true
	}
	if resp.Page == 1{
		resp.Pref = false
	}else {
		resp.Pref = true
	}
	fmt.Println("page",resp.Page,"pages",pages)
	resp.Object = object
	resp.Ret = 200
}
