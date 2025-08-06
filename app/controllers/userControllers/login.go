package usercontrollers

import (
	"CMS/app/models"
	"CMS/app/services/userServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	//接收参数
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//判断用户是否存在
	err = userServices.CheckUserExistByUsername(data.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200506, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	//获取用户信息
	var user *models.User
	user, err = userServices.GetUserByUsername(data.Username)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//判断密码是否正确

	flag := userServices.ComparePwd(data.Password, user.Password)
	if !flag {
		utils.JsonErrorResponse(c, 200507, "密码错误")
		return
	}
	utils.JsonSuccessResponse(c, user)

}
