package studentcontrollers

import (
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteData struct {
	UserID int `form:"user_id" json:"user_id" binding:"required"`
	PostID int `form:"post_id" json:"id" binding:"required"`
}

func Delete(c *gin.Context){
	var data DeleteData
	err:=c.ShouldBindQuery(&data)//绑定请求的参数到结构体中
	if err!=nil{
		utils.JsonErrorResponse(c,200501,"参数错误")
		return
	}

	err=studentservices.Delete(data.PostID)//删除某一条帖子
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