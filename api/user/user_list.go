package user

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/common/ctype"
	"gvb_server/models/common/response"
	"gvb_server/service/common"
	"gvb_server/utils/desens"
	"gvb_server/utils/jwt"
)

func (this *UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var page models.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.CommonList(models.UserModel{}, common.Option{
		Page:  page,
		Debug: true,
	})

	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 不是管理员，隐藏用户名
			user.UserName = ""
		}
		// 不管是不是管理员，电话和邮箱都要脱敏
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)

		users = append(users, user)
	}
	response.OkWithList(users, count, c)
}
