package controllers

import (
	"sync"
	"time"
	"fmt"
	"msw_bg/models"
	"msw_bg/util"
	"strings"
	"msw_bg/service"
)

type DishController struct {
	BaseController
}

func (this *DishController)ToDish(){
	resp :=service.GetDishList(1,util.PAGE_SIZE,"")
	this.Data["object"] = resp
	fmt.Println(resp)
	this.Data["url"] = "/dish/pagedish"
	this.IsneedTemplate()
	this.TplName = "dish.html"
	return
}

func (this *DishController)PageDish(){
	var resp models.PageResp
	resp.Ret = 403
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	name:= this.GetString("name")
	page,_ := this.GetInt("page")
	resp = service.GetDishList(page,util.PAGE_SIZE,name)
}


func (this *DishController)CreateDish(){
	var resp models.CreateDishResp
	resp.Ret = 403
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	if this.User == nil {
		resp.Msg = "请先登入"
	}
	s,f,err := this.GetFile("img")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	id := int(time.Now().Unix())
	path := util.SaveImg(s,f,id,"dish_img/")
	path = "http://localhost:8080/static/img/dish_img/"+path
	dishname := this.GetString("dishname")
	describe := this.GetString("describe")
	resp = service.CreateNewDish(this.User.Id,dishname,path,describe)
}

func (this *DishController)Delete(){
	var resp models.BaseResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	id,err := this.GetInt("dishid")
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	err = models.DeleteDish(id)
	if err != nil {
		resp.Msg = err.Error()
		return
	}
	resp.Ret = 200
}

func (this *DishController)UpdateDish(){
	var resp models.BaseResp
	resp.Ret = 403
	defer func (){
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	main := this.GetString("main")
	second := this.GetString("second")
	taste := this.GetString("taste")
	system := this.GetString("system")
	dish,_:=this.GetInt("dishid")
	resp = service.UpdateDishDetail(dish,taste,system,main,second)
}

func (this *DishController)UpdateDishStep(){
	var resp models.BaseResp
	defer func() {
		this.Data["json"] = resp
		this.ServeJSON()
	}()
	step,_ := this.GetInt("step")
	desc := this.GetString("describe")
	dishid,_ := this.GetInt("dishid")
	str := strings.Split(desc,"/")
	fmt.Println(str)
	j := step-1
	fmt.Println("j",j)
	for i:=step;i>= 1;i--{
		if  j<0 {
			break;
		}
		err := models.InsertDishStep(dishid,i,str[j])
		fmt.Println("i:",i,"str",str[j])
		if err != nil {
			resp.Msg = err.Error()
			return
		}
		j--
	}
	resp.Ret = 200
}

var l sync.RWMutex

var Step int

func (this *DishController)GetStepImg(){
	fmt.Println("开始访问了")
	s,f,err := this.GetFile("file")
	if err != nil {
		return
	}
	fmt.Println("开始访问了")
	//该方法只能被一个线程访问
	l.Lock()
	id := int(time.Now().Unix())
	time.Sleep(1*time.Second)
	fmt.Println("过一秒")
	path := util.SaveImg(s,f,id,"step/")
	path = "http://127.0.0.1:8080/static/img/step/"+path
	ids,_ := models.GetLastDish()
	Step,_ = models.CountStep(ids)
	models.InsertImgstep(ids,path,Step)
	imgs,_ := models.GetImg(ids)
	fmt.Println("存入的图片长度",len(imgs))
	rightposition := util.GetRegisterImgPosition(imgs)
	length := len(rightposition)
	count := models.GetCountImg(ids)
	fmt.Println("count",count,"length",length)
	if count == length {
		j := 0
		for i :=1 ;i <= count ;i++{
			models.UpdateRight(ids,rightposition[j],i)
			j++
		}
	}
	l.Unlock()
	this.Data["ok"] = "ok"
	this.ServeJSON()
}