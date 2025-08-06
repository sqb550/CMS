package userServices

import (
	"CMS/app/models"
	"CMS/config/database"
	"unicode"
)

func CheckUserExistByUsername(Username string) error {
	result := database.DB.Where("username=?", Username).First(&models.User{})
	return result.Error
}

func GetUserByUsername(Username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username=?", Username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2

}

func Register(user models.User) error {
	result := database.DB.Create(&user)
	return result.Error
}

func IsDigitsOnly(username string) bool {
	if username == "" {
		return false
	}
	for _, r := range username {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func IsLengthValid(password string)bool{
	length:=len(password)
	if length>=8 && length<=16{
		return true
	}
	return false

}

func IsUserTypeVaild(user_type int)bool{
	if user_type==1 || user_type==2{
		return true
	}
	
	return false
}