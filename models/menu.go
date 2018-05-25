package models

import "github.com/astaxie/beego/orm"

type Menu struct {
	Id           int
	Uid          int    `description:"用户id"`
	MenuName     string `description:"菜单名"`
	MenuImg      string `description:"菜单图片"`
	CreateDate   string `description:"创建日期"`
	MenuDescribe string `description:"菜单描述"`
	CollectCount int    `description:"收藏人数"`
	PopularCount int    `description:"人气"`
	Counts        int    `description:"菜谱的总数"`
}

//官方uid = 0
//新建菜单
func InsertMenu(uid int,menuname string,menuimg string,describe string)(err error){
	sql := `INSERT INTO menu (uid,menu_name,menu_img,menu_describe,create_date) VALUES (?,?,?,?,NOW())`
	o := orm.NewOrm()
	_,err = o.Raw(sql,uid,menuname,menuimg,describe).Exec()
	return
}


//检查菜单名
func CheckMenuName(uid int,menuname string)(count int,err error){
	sql := `SELECT COUNT(1) FROM menu WHERE uid = ? AND menu_name = ?`
	o := orm.NewOrm()
	err = o.Raw(sql,uid,menuname).QueryRow(&count)
	return
}


