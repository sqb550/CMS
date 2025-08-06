package studentcontrollers

import (
	"CMS/app/models"
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
)
type List struct{
	PostList []models.Post `json:"post_list"`
}

func Show(c *gin.Context){
	result,err:=studentservices.ShowPost()//result为post中的结构体数组
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
	}
	var results List
	results.PostList=result
	utils.JsonSuccessResponse(c,results)
}