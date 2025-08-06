package studentcontrollers

import (
	"CMS/app/models"
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdateData struct {
	ID      uint   `json:"post_id" binding:"required"`
	UserID  int    `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func Update(c *gin.Context){
	var data UpdateData
	err:=c.ShouldBindJSON(&data)//绑定
	if err!=nil{
		utils.JsonErrorResponse(c,200501,"参数错误")
		return
	}

	err=studentservices.Update(models.Post{
		ID : data.ID,
		UserID: data.UserID,
		Content: data.Content,

	})//更新数据
	if err!=nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200510, "该帖子不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return 
	}
	utils.JsonSuccessResponse(c,nil)

}