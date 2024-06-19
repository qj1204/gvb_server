package user

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/response"
	"gvb_server/service/common_service"
	"gvb_server/utils/desens"
	"gvb_server/utils/jwt"
)

// UserListView 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Param data query models.Page  false  "查询参数"
// @Param token header string  true  "token"
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.UserModel]}
func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var page models.Page
	if err := c.ShouldBindQuery(&page); err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common_service.CommonList(models.UserModel{}, common_service.Option{
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
