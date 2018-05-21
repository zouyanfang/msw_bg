package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"msw_bg/util"
	"fmt"
)

func init(){
	err :=orm.RegisterDataBase("default","mysql",util.MYSQL_URL)
	if err != nil {
		fmt.Print(err.Error())
	}
}
