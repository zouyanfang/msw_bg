package util

import (
	"github.com/astaxie/beego"
	"fmt"
)

var (
	MYSQL_URL string
	PAGE_SIZE = 4
)

func init() {
	runmode := beego.AppConfig.String("runmode")
	config, err := beego.AppConfig.GetSection(runmode)
	if err != nil {
		fmt.Println(err.Error())
		panic("配置文件初始化失败")
	}
	MYSQL_URL = config["mysql_url"]
}
