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

