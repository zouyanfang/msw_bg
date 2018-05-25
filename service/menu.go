package service

import (
	"msw_bg/models"

)

//新建菜单
func CreateNewMenu(uid int, menuname string,menuimg string,describe string)(resp models.BaseResp){
	count,err := models.CheckMenuName(uid,menuname)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	if count != 0 {
		resp.Msg = "菜单重名了！请换一个名字"
	}
	err = models.InsertMenu(uid,menuname,menuimg,describe)
	if err != nil {
		resp.Msg = "创建失败"+err.Error()
		return
	}
	resp.Ret = 200
	return
}
