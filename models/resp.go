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

type CountResp struct {
	BaseResp
	TotalNum     int
	YesterdayNum int
	TodayNum     int
}

type PageResp struct {
	BaseResp
	Page int
	Count int
	Next bool
	Pref bool

}
