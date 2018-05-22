package controllers

import "fmt"

type UserController struct {
	BaseController
}

func (this *UserController)ToUser(){
	fmt.Println("22")
	this.IsneedTemplate()
	this.TplName = "user.html"
	return
}
