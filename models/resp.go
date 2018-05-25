package models

type BaseResp struct {
	Ret    int
	Msg    string
	Object interface{}
}

type CreateDishResp struct {
	BaseResp
	Id int
}

type RegisterCountResp struct {
	BaseResp
	TotalNum     int //总的注册人数
	YesterdayNum int //昨日注册人数
	TodayNum     int //今日注册人数
}
