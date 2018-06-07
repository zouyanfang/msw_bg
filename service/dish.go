package service

import (
	"msw_bg/models"
	"msw_bg/util"
)

func CreateNewDish(uid int,dishname string,dishimg string,describe string)(resp models.CreateDishResp){
	err := models.CreateDish(uid,dishname,dishimg,describe)
	if err != nil {
		resp.Msg = err.Error()
	}
	id,_ := models.GetLastDish()
	resp.Id = id
	resp.Ret = 200
	return
}

func UpdateDishDetail(dishid int,taste,system string,main ,second string)(resp models.BaseResp){
	err := models.UpdateDish(dishid,taste,system,main,second)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
	return
}

func GetDishList(page,pageSize int,name string)(resp models.PageResp){
	if name != ""{
		name = " '%"+name+"%' "
	}
	start := util.StartIndex(page,pageSize)
	list,err := models.GetDishList(start,pageSize,name)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	count,err := models.GetDishCount(name)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Page = page
	ParsePage(&resp,count,pageSize,list)
	return
}

