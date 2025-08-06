package usercontrollers

import (
	"CMS/app/models"
	"CMS/app/services/userServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterDate struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType int    `json:"user_type" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func Register(c *gin.Context) {
	var data RegisterDate
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//判断用户名是否为纯数字
	flag:=userServices.IsDigitsOnly(data.Username)
	if !flag{
		utils.JsonErrorResponse(c,200502,"用户名必须为纯数字")
		return
	}
	//判断密码长度是否符合要求
	flag=userServices.IsLengthValid(data.Password)
	if !flag{
		utils.JsonErrorResponse(c,200503,"密码长度必须在8-16位")
		return
	}
	//判断用户类型是否正确

	flag=userServices.IsUserTypeVaild(data.UserType)
	if !flag{
		utils.JsonErrorResponse(c,200504,"用户类型错误")
		return
	}

	//判断该用户名是否存在
	err = userServices.CheckUserExistByUsername(data.Username)
	if err == nil {
		utils.JsonErrorResponse(c, 200505, "用户名已存在")
		return
	} else if err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	err = userServices.Register(models.User{
		Username: data.Username,
		Name:     data.Name,
		Password: data.Password,
		UserType: data.UserType,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
