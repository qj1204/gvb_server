package data_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"time"
)

type DateCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type DateCountResponse struct {
	DateList  []string `json:"date_list"`
	LoginData []int    `json:"login_data"`
	SignData  []int    `json:"sign_data"`
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
// @Success 200 {object} res.Response{data=[]DateCountResponse}
func (DataApi) DataLoginView(c *gin.Context) {
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

	var response DateCountResponse

	var dateTypeNum = dateMap[cr.Date]
	query.Where(fmt.Sprintf("date_sub(curdate(), interval %d day) <= created_at", dateTypeNum))
	aDay := now.AddDate(0, 0, -dateTypeNum)
	for i := 1; i <= dateTypeNum; i++ {
		response.DateList = append(response.DateList, aDay.AddDate(0, 0, i).Format("2006-01-02"))
	}

	type dateCountType struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}

	var dateCountList []dateCountType
	global.DB.Model(models.LoginDataModel{}).Where(query).
		Select(
			"date_format(created_at, '%Y-%m-%d') as date",
			"count(id) as count").
		Group("date").Scan(&dateCountList)
	var dateLoginCountMap = map[string]int{}
	for _, countType := range dateCountList {
		dateLoginCountMap[countType.Date] = countType.Count
	}
	for _, s := range response.DateList {
		count, _ := dateLoginCountMap[s]
		response.LoginData = append(response.LoginData, count)
	}

	var signCounList []dateCountType
	global.DB.Model(models.UserModel{}).Where(query).
		Select(
			"date_format(created_at, '%Y-%m-%d') as date",
			"count(id) as count").
		Group("date").Scan(&signCounList)
	var dateSignCountMap = map[string]int{}
	for _, countType := range signCounList {
		dateSignCountMap[countType.Date] = countType.Count
	}
	for _, s := range response.DateList {
		count, _ := dateSignCountMap[s]
		response.SignData = append(response.SignData, count)
	}

	res.OkWithData(response, c)

}
