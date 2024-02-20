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
		response.FailWithCode(response.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.CommonList(models.UserModel{}, common.Option{
		Page: page,
	})

	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 不是管理员
			user.UserName = ""
		}
		// 脱敏
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)

		users = append(users, user)
	}
	response.OkWithList(users, count, c)
}
