package usercontrollers

import (
	"unicode"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	apiexception "CMS/app/apiException"
	"CMS/app/models"
	"CMS/app/services/userServices"
	"CMS/app/utils"
)

type RegisterData struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType int    `json:"user_type" binding:"required"`
	Username string `json:"username" binding:"required"`
}

func Register(c *gin.Context) {
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ParamError, err)
		return
	}
	//判断用户名是否为纯数字
	flag := true
	if data.Username == "" {
		flag = false
	}
	for _, r := range data.Username {
		if !unicode.IsDigit(r) {
			flag = false
		}
	}
	if !flag {
		apiexception.AbortWithException(c, apiexception.UsernameError, err)
	}

	//判断密码长度是否符合要求
	flag = false
	length := len(data.Password)
	if length >= 8 && length <= 16 {
		flag = true
	}

	if !flag {
		apiexception.AbortWithException(c, apiexception.PasswordLengthError, err)
		return
	}
	//判断用户类型是否正确

	flag = false

	if data.UserType == 1 || data.UserType == 2 {
		flag = true
	}

	if !flag {
		apiexception.AbortWithException(c, apiexception.UserTypeError, err)
		return
	}

	//判断该用户名是否存在
	result, err := userServices.GetUserByUsername(data.Username)
	if result != nil {
		apiexception.AbortWithException(c, apiexception.UserExist, err)
		return
	} else if err != gorm.ErrRecordNotFound {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	err = userServices.Register(&models.User{
		Username: data.Username,
		Name:     data.Name,
		Password: string(hashPassword),
		UserType: data.UserType,
	})
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
