package models

type BaseResp struct {
	Ret int
	Msg string
	Object interface{}
}

type CreateDishResp struct {
	BaseResp
	Id int
}

//分页返回体
type PageResp struct {
	BaseResp
	Page int
	Count int
	Next bool
	Pref bool
}
