package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Admin struct {
	Id int
	Account string
	Pwd string
	Role int
	State int
	CreateTime string
}

func Login(account string,pwd string)(a *Admin){
	sql := `SELECT * FROM admin WHERE account = ? AND pwd = ?  `
	o := orm.NewOrm()
	err := o.Raw(sql,account,pwd).QueryRow(&a)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func GetAdminList(pageIndex,pageSize int)(list []Admin,err error){
	sql := `SELECT * FROM admin WHERE role = 0 LIMIT ?,?`
	_,err = orm.NewOrm().Raw(sql,pageIndex,pageSize).QueryRows(&list)
	return
}

func CountAdmin()(count int,err error){
	sql := `SELECT COUNT(1) FROM admin`
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

func FindAccount(account string)(count int,err error){
	sql := `SELECT COUNT(1) FROM admin WHERE account = ?`
	err = orm.NewOrm().Raw(sql,account).QueryRow(&count)
	return
}

func InsertAdmin(account ,pwd string)(err error){
	sql := `INSERT INTO admin (account,pwd,role,create_time,state) VALUES (?,?,0,NOW(),1)`
	_,err = orm.NewOrm().Raw(sql,account,pwd).Exec()
	return
}