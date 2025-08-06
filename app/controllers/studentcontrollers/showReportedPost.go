package studentcontrollers

import (
	
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
)

type ReportedPostData struct {
	UserID int `form:"user_id" json:"user_id"`
}

type ShowReportedPostData struct{
	PostID int `json:"post_id"`
	Content string `json:"content"`
	Reason string `json:"reason"`
	Status int `json:"status"`
}

type ReportedPostList struct{
	ReportList []ShowReportedPostData `json:"report_list"`
}


func ShowReportedPost(c *gin.Context) {
	var data ReportedPostData
	err := c.ShouldBindQuery(&data)
	if err!=nil{
		utils.JsonErrorResponse(c,200501,"参数错误")
		return
	}
	ReportedPost,err:=studentservices.ShowReportedPost(data.UserID)//获取举报人id为data，userid所举报的所有帖子
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	
	var NewReportedPost []ShowReportedPostData
	for _,data:=range ReportedPost{
		NewReportedPost=append(NewReportedPost,
		ShowReportedPostData{
			PostID:data.PostID,
			Content: data.Content,
			Reason: data.Reason,
			Status: data.Status,
		})
	}
	var result ReportedPostList
	result.ReportList=NewReportedPost
	utils.JsonSuccessResponse(c,result)//返回举报的所有帖子

}