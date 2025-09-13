package studentControllers

import (
	apiexception "CMS/app/apiException"
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
)

type PostsData struct {
	PostID int `json:"post_id"`
}

// 点赞处理函数
func LikePost(c *gin.Context) {
	var data PostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ParamError, err)
		return
	}

	_, ok, err := utils.GetPostFromCache(uint(data.PostID),c)
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
		return
	}
	if !ok {
		Post, err := studentservices.GetPost(data.PostID)
		if err != nil {
			apiexception.AbortWithException(c, apiexception.ServerError, err)
			return
		}

		err = utils.SetPostToCache(Post,c)
		if err != nil {
			apiexception.AbortWithException(c, apiexception.ServerError, err)
			return
		}
	}
	err=utils.LikesIncr(data.PostID,c)
	if err != nil {
			apiexception.AbortWithException(c, apiexception.ServerError, err)
			return
		}
	

	utils.JsonSuccessResponse(c, nil)
}
