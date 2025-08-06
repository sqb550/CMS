package studentcontrollers

import (
	"CMS/app/models"
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
)

type ReleaseData struct{
	Content string `json:"content" binding:"required"`
	UserID int `json:"user_id" binding:"required"`
}

func Release(c *gin.Context){
	var data ReleaseData
	err:=c.ShouldBindJSON(&data)
	if err!=nil{
		utils.JsonErrorResponse(c,200501,"参数错误")
	}
	//将帖子信息添加到数据表post
	err = studentservices.ReleasePost(models.Post{
		Content: data.Content,
		UserID:     data.UserID,
	
		
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}