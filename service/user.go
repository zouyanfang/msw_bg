package service

import "msw_bg/models"

func GetUserRegisterNum() (resp models.RegisterCountResp) {
	resp.Ret = 403
	//获取注册总人数
	totalNum,err := models.GetTotalRegisterNum()
	if err != nil {
		resp.Msg = "获取注册总人数失败"
		return
	}
	resp.TotalNum = totalNum
	//获取今日注册人数
	todayNum,err := models.GetTodayRegNum()
	if err != nil {
		resp.Msg = "获取今日注册人数失败"
		return
	}
	resp.TodayNum = todayNum
	//获取昨日注册人数
	yesNum,err := models.GetYesRegNum()
	if err != nil {
		resp.Msg = "获取昨日注册人数失败"
		return
	}
	resp.YesterdayNum = yesNum
	resp.Ret = 200
	return
}
