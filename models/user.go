package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

//获取总注册人数
func GetTotalRegisterNum() (totalNum int,err error) {
	sql := `SELECT COUNT(1) FROM users`
	err = orm.NewOrm().Raw(sql).QueryRow(&totalNum)
	fmt.Println("qweqweqw",totalNum)
	return 
}

//获取今日注册人数
func GetTodayRegNum() (todayNum int,err error) {
	sql := `SELECT COUNT(1) FROM users WHERE register_date = curdate()`
	err = orm.NewOrm().Raw(sql).QueryRow(&todayNum)
	return
}

//获取昨日注册人数
func GetYesRegNum() (yesNum int,err error) {
	sql := `SELECT COUNT(1) FROM users WHERE register_date = curdate()-1`
	err = orm.NewOrm().Raw(sql).QueryRow(&yesNum)
	return
}

//获取总留言数
func GetTotalMessage() (totalNum int,err error) {
	sql := `SELECT COUNT(1) FROM user_message`
	err = orm.NewOrm().Raw(sql).QueryRow(&totalNum)
	return
}

//获取昨日留言数
func GetYesMessage() (yesNum int,err error)  {
	sql := `SELECT COUNT(1) FROM user_message WHERE DATE(create_time) = CURDATE() - 1`
	err = orm.NewOrm().Raw(sql).QueryRow(&yesNum)
	return
}

//获取今日留言数
func GetTodayMessage() (todayNum int,err error) {
	sql := `SELECT COUNT(1) FROM user_message WHERE DATE(create_time) = CURDATE()`
	err = orm.NewOrm().Raw(sql).QueryRow(&todayNum)
	return
}

type User struct {
	Id           int
	Name         string `description:"名字"`
	Account      string `description:"账号"`
	Pwd          string `description:"密码"`
	UserImg      string `description:"头像"`
	RegisterDate string `description:"注册日期"`
}

//留言板
type MessageBoard struct {
	Id         int
	Title      string `description:"留言标题"`
	Content    string `description:"留言内容"`
	CreateTime string `description:"创建时间"`
	Name       string `description:"用户昵称"`
	UserImg    string `description:"用户图片"`
}

//用户列表
func UserList(pageIndex,pageSize int,name string)(list []MessageBoard	,err error){
	sql := `SELECT m.id,m.content,create_time,u.name FROM user_message m LEFT JOIN users u ON m.uid = u.id  `
	if name != ""{
		sql += "where u.name like "+name
	}
	sql += ` ORDER BY create_time DESC LIMIT ?,?`
	_,err = orm.NewOrm().Raw(sql,pageIndex,pageSize).QueryRows(&list)
	fmt.Println(sql)
	return
}

//用户总数
func UserCount(uid int)(count int,err error)  {
	sql := `SELECT COUNT(1) FROM user_message `
	if uid != 0 {
		sql += " where uid = ?"
		err = orm.NewOrm().Raw(sql,uid).QueryRow(&count)
		return
	}
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	fmt.Println(sql,uid)
	return
}


func GetNewsList(pageIndex,pageSize int,name string)(list []MessageBoard,err error){
	sql := `SELECT title,id,create_time,content FROM message_board `
	if name != ""{
		sql += " WHERE title like "+name
	}
	sql += " ORDER BY create_time DESC LIMIT ?,? "
	_,err = orm.NewOrm().Raw(sql,pageIndex,pageSize).QueryRows(&list)
	fmt.Println(list)
	return
}

func CountNews(name string)(count int,err error){
	sql := `SELECT COUNT(1) FROM message_board`
	if name != "" {
		sql += " WHERE title like "+name
		err = orm.NewOrm().Raw(sql).QueryRow(&count)
		return
	}
	err = orm.NewOrm().Raw(sql).QueryRow(&count)
	return
}

func DeleteNews(id int)(err error){
	sql := `DELETE FROM message_board WHERE id = ?`
	_,err = orm.NewOrm().Raw(sql,id).Exec()
	return
}

func DeleteAdmin(id int)(err error){
	sql := `DELETE FROM admin where id = ?`
	_,err = orm.NewOrm().Raw(sql,id).Exec()
	return
}

func UpdateSysUser(uid int,state int) error {
	sql := `UPDATE admin set state = ? WHERE id = ?`
	_,err := orm.NewOrm().Raw(sql,state,uid).Exec()
	fmt.Println(sql,uid,state)
	return err
}

//获取用户detail
func GetUserMsgDetail(id int)(m MessageBoard,err error){
	sql := `SELECT * FROM user_message WHERE id = ?`
	err = orm.NewOrm().Raw(sql,id).QueryRow(&m)
	return
}

func FindByUid(name string)(err error,id int){
	sql := `SELECT id FROM users WHERE name = ?`
	err = orm.NewOrm().Raw(sql,name).QueryRow(&id)

	return
}