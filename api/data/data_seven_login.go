package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/common/response"
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

func (this *DataApi) SevenLoginView(c *gin.Context) {

	var signDateCount, loginDateCount []DateCount
	// global.DB.Raw("select DATE_FORMAT(created_at,'%Y-%m-%d') as date, count(id) as count from login_data_models where DATE_SUB(CURDATE(), INTERVAL 7 day) <= created_at GROUP BY date").Scan(&dateCount)
	global.DB.Model(models.LoginDataModel{}).
		Where("DATE_SUB(CURDATE(), INTERVAL 7 day) <= created_at").
		Select("DATE_FORMAT(created_at,'%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&loginDateCount)

	global.DB.Model(models.UserModel{}).
		Where("DATE_SUB(CURDATE(), INTERVAL 7 day) <= created_at").
		Select("DATE_FORMAT(created_at,'%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&signDateCount)

	var signDateCountMap = make(map[string]int)
	var loginDateCountMap = make(map[string]int)
	for _, v := range signDateCount {
		signDateCountMap[v.Date] = v.Count
	}
	for _, v := range loginDateCount {
		loginDateCountMap[v.Date] = v.Count
	}

	var dateList []string
	var signCountList []int
	var loginCountList []int
	now := time.Now()
	for i := -6; i <= 0; i++ {
		day := now.AddDate(0, 0, i).Format("2006-01-02")
		fmt.Println(day)
		dateList = append(dateList, day)
		signCountList = append(signCountList, signDateCountMap[day])
		loginCountList = append(loginCountList, loginDateCountMap[day])
	}
	response.OkWithData(DateCountResponse{
		DateList:       dateList,
		SignCountList:  signCountList,
		LoginCountList: loginCountList,
	}, c)
}
