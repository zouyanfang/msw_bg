package controllers

type NewsController struct {
	BaseController
}

func (this *NewsController)ToNews(){
	this.IsneedTemplate()
	this.TplName = "news.html"
}
