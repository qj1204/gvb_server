package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"time"
)

type DateCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type DateCountResponse struct {
	DateList       []string `json:"date_list"`
	SignCountList  []int    `json:"sign_count_list"`
	LoginCountList []int    `json:"login_count_list"`
}

type LoginRequest struct {
	Date int `json:"date" form:"date"` // 1 七天 2 一个月 3  两个月 4 三个月 5 六个月  6 一年
}

// DataLoginView 七日登录，注册数据
// @Tags 数据管理
// @Summary 七日登录，注册数据
// @Description 七日登录，注册数据
// @Router /api/data_login [get]
// @Produce json
// @Success 200 {object} response.Response{data=[]DateCountResponse}
func (DataApi) DataLoginView(c *gin.Context) {

	//var signDateCount, loginDateCount []DateCount
	//// global.DB.Raw("select DATE_FORMAT(created_at,'%Y-%m-%d') as date, count(id) as count from login_data_models where DATE_SUB(CURDATE(), INTERVAL 7 day) <= created_at GROUP BY date").Scan(&dateCount)
	//global.DB.Model(models.LoginDataModel{}).
	//	Where("DATE_SUB(CURDATE(), INTERVAL 7 day) <= created_at").
	//	Select("DATE_FORMAT(created_at,'%Y-%m-%d') as date", "count(id) as count").
	//	Group("date").
	//	Scan(&loginDateCount)
	//
	//global.DB.Model(models.UserModel{}).
	//	Where("DATE_SUB(CURDATE(), INTERVAL 7 day) <= created_at").
	//	Select("DATE_FORMAT(created_at,'%Y-%m-%d') as date", "count(id) as count").
	//	Group("date").
	//	Scan(&signDateCount)
	//
	//var signDateCountMap = make(map[string]int)
	//var loginDateCountMap = make(map[string]int)
	//for _, v := range signDateCount {
	//	signDateCountMap[v.Date] = v.Count
	//}
	//for _, v := range loginDateCount {
	//	loginDateCountMap[v.Date] = v.Count
	//}
	//
	//var dateList []string
	//var signCountList []int
	//var loginCountList []int
	//now := time.Now()
	//for i := -6; i <= 0; i++ {
	//	day := now.AddDate(0, 0, i).Format("2006-01-02")
	//	fmt.Println(day)
	//	dateList = append(dateList, day)
	//	signCountList = append(signCountList, signDateCountMap[day])
	//	loginCountList = append(loginCountList, loginDateCountMap[day])
	//}
	//response.OkWithData(DateCountResponse{
	//	DateList:       dateList,
	//	SignCountList:  signCountList,
	//	LoginCountList: loginCountList,
	//}, c)

	var cr LoginRequest
	_ = c.ShouldBindQuery(&cr)

	var query = global.DB.Where("")
	now := time.Now()

	var dateMap = map[int]int{
		0: 7,
		1: 7,
		2: 30,
		3: 60,
		4: 90,
		5: 180,
		6: 365,
	}

	var res DateCountResponse

	var dateTypeNum = dateMap[cr.Date]
	query.Where(fmt.Sprintf("date_sub(curdate(), interval %d day) <= created_at", dateTypeNum))
	aDay := now.AddDate(0, 0, -dateTypeNum)
	for i := 1; i <= dateTypeNum; i++ {
		res.DateList = append(res.DateList, aDay.AddDate(0, 0, i).Format("2006-01-02"))
	}

	var dateCountList []DateCount
	global.DB.Model(models.LoginDataModel{}).Where(query).Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").Scan(&dateCountList)
	var dateLoginCountMap = map[string]int{}
	for _, dataCount := range dateCountList {
		dateLoginCountMap[dataCount.Date] = dataCount.Count
	}
	for _, s := range res.DateList {
		count, _ := dateLoginCountMap[s]
		res.LoginCountList = append(res.LoginCountList, count)
	}

	var signCountList []DateCount
	global.DB.Model(models.UserModel{}).Where(query).Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").Scan(&signCountList)
	var dateSignCountMap = map[string]int{}
	for _, dataCount := range signCountList {
		dateSignCountMap[dataCount.Date] = dataCount.Count
	}
	for _, s := range res.DateList {
		count, _ := dateSignCountMap[s]
		res.SignCountList = append(res.SignCountList, count)
	}

	response.OkWithData(res, c)
}
