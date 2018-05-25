package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

//获取总注册人数
func GetTotalRegisterNum() (totalNum int,err error) {
	sql := `SELECT COUNT(1) FROM users`
	err = orm.NewOrm().Raw(sql).QueryRow(&totalNum)
	fmt.Println("qweqweqw",totalNum)
	return 
}

//获取今日注册人数
func GetTodayRegNum() (todayNum int,err error) {
	sql := `SELECT COUNT(1) FROM users WHERE register_date = curdate()`
	err = orm.NewOrm().Raw(sql).QueryRow(&todayNum)
	return
}

//获取昨日注册人数
func GetYesRegNum() (yesNum int,err error) {
	sql := `SELECT COUNT(1) FROM users WHERE register_date = curdate()-1`
	err = orm.NewOrm().Raw(sql).QueryRow(&yesNum)
	return
}