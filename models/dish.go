package models

import (
	"github.com/astaxie/beego/orm"
)

func GetLastDish()(id int,err error){
	sql := `SELECT id FROM dish ORDER BY id DESC LIMIT 0,1`
	o := orm.NewOrm()
	err = o.Raw(sql).QueryRow(&id)
	return
}


func CountStep(dishid int,)(count int,err error){
	sql := `SELECT COUNT(1) FROM dish_step  WHERE dish_id = ? AND step_img is NULL `
	o := orm.NewOrm()
	err = o.Raw(sql,dishid).QueryRow(&count)
	return
}

func InsertImgstep(dishid int,stepimg string,step int)(err error){
	sql := `UPDATE dish_step SET step_img = ? WHERE dish_id = ? AND step = ? `
	o := orm.NewOrm()
	_,err = o.Raw(sql,stepimg,dishid,step).Exec()
	return
}

func GetImg(dishId int)(StepImg []string ,err error){
	sql := `SELECT step_img FROM dish_step WHERE dish_id = ? AND step_img is not NULL`
	o := orm.NewOrm()
	o.Raw(sql,dishId).QueryRows(&StepImg)
	return
}

func GetCountImg(dishId int)(count int){
	sql := `SELECT COUNT(1) FROM dish_step WHERE dish_id = ?`
	o := orm.NewOrm()
	o.Raw(sql,dishId).QueryRow(&count)
	return
}


func UpdateRight(dishid int,img string,step int)(){
	sql := `UPDATE dish_step SET step_img =? WHERE dish_id = ? AND step = ?`
	o := orm.NewOrm()
	o.Raw(sql,img,dishid,step).Exec()
	return
}

func InsertDishStep(dishid ,step int,stepdescribe string)(err error){
	sql := `INSERT INTO dish_step (dish_id,step,step_describe) VALUES (?,?,?)`
	o:= orm.NewOrm()
	_,err = o.Raw(sql,dishid,step,stepdescribe).Exec()
	return
}

func CreateDish(uid int,dishname,dishimg,describe string)(err error){
	sql := `INSERT INTO dish (uid,dish_name,dish_img,release_date,release_role,dish_describe) VALUES (?,?,?,NOW(),2,?)`
	o := orm.NewOrm()
	_,err = o.Raw(sql,0,dishname,dishimg,describe).Exec()
	return
}

func UpdateDish(dishid int,taste,system string,main ,second string)(err error){
	sql := `UPDATE dish SET main_material = ?,second_material = ?,tasty = ?,dish_system = ? WHERE id = ?`
	o := orm.NewOrm()
	_,err = o.Raw(sql,main,second,taste,system,dishid).Exec()
	return
}
