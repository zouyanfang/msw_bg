package models

import (
	"github.com/astaxie/beego/orm"
)

//菜谱
type Dish struct {
	Id             int
	Uid            int    `description:"用户id"`
	DishName       string `description:"菜名"`
	DishImg        string `description:"菜成品图"`
	IsSlideshow    int    `description:"是否是轮播图"`
	ReleaseDate    string `description:"发布时间"`
	MainMaterial   string `description:"主料"`
	SecondMaterial string `description:"辅料"`
	ReleaseRole    int    `description:"发布角色 0 官方 1用户"`
	Tasty          string `description:"口味"`
	DishSystem     string `description:"菜系"`
	DishDescribe   string `description:"菜谱描述"`
	CollectCount   int    `description:"收藏人数"`
	PopularCount   int    `description:"人气"`
}

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

func GetDishList(pageIndex,pageSize int,name string)(list []Dish,err error){
	sql := `SELECT id,dish_name,popular_count,collect_count FROM dish `
	if name != ""{
		sql += " WHERE dish_name like "+name
	}
	sql += " LIMIT ?,? "
	_,err = orm.NewOrm().Raw(sql,pageIndex,pageSize).QueryRows(&list)
	return
}

func GetDishCount(name string )(count int,err error){
	sql := `SELECT count(1) FROM dish`
	if name != ""{
		sql += " WHERE dish_name like "+name
		err = orm.NewOrm().Raw(sql).QueryRow(&count)
		return
	}
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

func ReleaseNews(uid int,title,content string)(err error ){
	sql := `INSERT INTO message_board (uid,title,content,create_time) VALUES (?,?,?,NOW())`
	_,err = orm.NewOrm().Raw(sql,uid,title,content).Exec()
	return
}


func DeleteDish(dishid int)(err error){
	o := orm.NewOrm()
	defer func() {
		if err != nil {
			o.Rollback()
			return
		}
		o.Commit()
	}()
	sql := `DELETE FROM dish WHERE id = ?`
	_,err = o.Raw(sql,dishid).Exec()
	if err != nil {
		return
	}
	sql = `DELETE FROM menu_dish WHERE dish_id = ?`
	_,err = o.Raw(sql,dishid).Exec()
	if err != nil {
		return
	}
	sql = `DELETE FROM dish_step WHERE dish_id = ?`
	_,err = o.Raw(sql,dishid).Exec()
	if err != nil {
		return
	}
	sql = `DELETE FROM user_collection WHERE dish_id = ?`
	_,err = o.Raw(sql,dishid).Exec()
	if err != nil {
		return
	}
	sql = `DELETE FROM dish_comment WHERE dish_id = ?`
	_,err = o.Raw(sql,dishid).Exec()
	if err != nil {
		return
	}
	return
}
