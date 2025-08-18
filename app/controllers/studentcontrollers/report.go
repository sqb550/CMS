package studentControllers

import (
	"github.com/gin-gonic/gin"

	apiexception "CMS/app/apiException"
	"CMS/app/models"
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"
)

type ReportData struct {
	PostID int    `json:"post_id" binding:"required"`
	UserID int    `json:"user_id" binding:"required"`
	Reason string `json:"reason" binding:"required"`
}

func Report(c *gin.Context) {

	var data ReportData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ParamError, err)
		return
	}

	PostResult, err := studentservices.GetPost(int(data.PostID))
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
		return
	} //获取被举报帖子的内容和发布者

	result, err := studentservices.SeekUser(PostResult.UserID) //寻找被举报人的username
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
	}
	err = studentservices.AddReportedPost(&models.ReportedPost{
		UserID:   data.UserID,
		PostID:   data.PostID,
		Content:  PostResult.Content,
		Reason:   data.Reason,
		Status:   0,
		Username: result.Username,
	}) //将数据内容添加到reportedpost数据表中
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
