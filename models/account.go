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
	Status int
}

func Login(account string,pwd string)(a *Admin){
	sql := `SELECT * FROM admin WHERE account = ? AND pwd = ?`
	o := orm.NewOrm()
	err := o.Raw(sql,account,pwd).QueryRow(&a)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
