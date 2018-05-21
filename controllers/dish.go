package controllers


type DishController struct {
	BaseController
}

func (this *DishController)ToDish(){
	this.IsneedTemplate()
	this.TplName = "dish.html"
	return
}