package apiexception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Code int
	Msg  string
}

// Error implements error.
func (e *Error) Error() string {
	panic("unimplemented")
}

var (
	ServerError         = NewError(200500, "系统异常，请稍后重试")
	ParamError          = NewError(200501, "参数错误")
	UsernameError       = NewError(200502, "用户名必须为纯数字")
	PasswordLengthError = NewError(200503, "密码长度必须在8-16位")
	UserTypeError       = NewError(200504, "用户类型错误")
	UserExist           = NewError(200505, "用户名已存在")
	UserNotFound        = NewError(200506, "用户名不存在")
	NotManagerError     = NewError(200507, "非管理员不具备审核权限")
	CheckError          = NewError(200508, "审核结果不正确")
	PostNotFound        = NewError(200509, "该帖子不存在")
	PasswordError       = NewError(200510, "密码错误")

	NotFound = NewError(200404, http.StatusText(http.StatusNotFound))
)

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func AbortWithException(c *gin.Context, apiError *Error, err error) {
	_ = c.AbortWithError(200, apiError)
}
