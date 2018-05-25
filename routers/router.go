package routers

import (
	"msw_bg/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/",&controllers.UserController{},"get:ToUser")
    beego.Router("/login",&controllers.AccountController{},"get:ToLogin")
    beego.AutoRouter(&controllers.UserController{})//用户controller
    beego.AutoRouter(&controllers.DishController{})//菜谱controller
    beego.AutoRouter(&controllers.MenuController{})//菜单controller
    beego.AutoRouter(&controllers.SysUserController{})//
    beego.AutoRouter(&controllers.AccountController{})
}
