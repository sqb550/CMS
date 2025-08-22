package managerControllers

import (
	"github.com/gin-gonic/gin"

	apiexception "CMS/app/apiException"
	managerservices "CMS/app/services/managerServices"
	"CMS/app/utils"
)

type ReportedPostData struct {
	ReportID int    `json:"report_id"`
	Username string `json:"username"`
	PostID   int    `json:"post_id"`
	Content  string `json:"content"`
	Reason   string `json:"reason"`
}

type ReportedPostList struct {
	ReportList []ReportedPostData `json:"report_list"`
}

func ShowReportedPosts(c *gin.Context) {

	ReportedPosts, err := managerservices.ReportedPostShow() //返回所有status为0的帖子
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
		return
	}
	var NewReportedPost []ReportedPostData
	for _, data := range ReportedPosts {
		NewReportedPost = append(NewReportedPost,
			ReportedPostData{
				ReportID: int(data.ID),
				Username: data.Username,
				PostID:   data.PostID,
				Content:  data.Content,
				Reason:   data.Reason,
			})
	}
	var result ReportedPostList
	result.ReportList = NewReportedPost
	utils.JsonSuccessResponse(c, result)

}
