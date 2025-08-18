package usercontrollers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	apiexception "CMS/app/apiException"
	"CMS/app/models"
	"CMS/app/services/userServices"
	"CMS/app/utils"
)

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ShowUser struct {
	UserID   int `json:"user_id"`
	UserType int `json:"user_type"`
}

func Login(c *gin.Context) {
	//接收参数
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ParamError, err)
		return
	}

	//获取用户信息和判断用户是否存在
	var user *models.User
	user, err = userServices.GetUserByUsername(data.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			apiexception.AbortWithException(c, apiexception.UserNotFound, err)
		} else {
			apiexception.AbortWithException(c, apiexception.ServerError, err)
		}
		return
	}

	//判断密码是否正确

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		apiexception.AbortWithException(c, apiexception.PasswordError, err)
		return
	}
	result := ShowUser{
		UserID:   int(user.ID),
		UserType: user.UserType,
	}
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()

	utils.JsonSuccessResponse(c, result)

}
