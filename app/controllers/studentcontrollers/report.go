package studentcontrollers

import (
	"CMS/app/models"
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
)

type ReportData struct {
	PostID    int   `json:"post_id" binding:"required"`
	UserID int    `json:"user_id" binding:"required"`
	Reason string `json:"reason" binding:"required"`
}

func Report(c *gin.Context){
	var data ReportData
	err:=c.ShouldBindJSON(&data)
	if err!=nil{
		utils.JsonErrorResponse(c,200501,"参数错误")
		return
	}

	PostResult,err:=studentservices.GetPost(int(data.PostID))
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}//获取被举报帖子的内容和发布者
	
	username,err:=studentservices.SeekUser(PostResult.UserID)//寻找被举报人的username
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
	}
	err=studentservices.AddReportedPost(models.ReportedPost{
		UserID:data.UserID,
		PostID: data.PostID,
		Content: PostResult.Content,
		Reason: data.Reason,
		Status: 0,
		Username: username,
	})//将数据内容添加到reportedpost数据表中
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	
	utils.JsonSuccessResponse(c,nil)
}