package studentControllers

import (
	apiexception "CMS/app/apiException"
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
)

type PostData struct {
	PostID int `form:"post_id" json:"post_id"`
}

func GetPostLikes(c *gin.Context) {
	var data PostData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ParamError, err)
		return
	}

	PostLikes, ok, err := utils.GetPostFromCache(uint(data.PostID),c)
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
		return
	}
	if ok {
		utils.JsonSuccessResponse(c, PostLikes)
		return
	}

	dbPost, err := studentservices.GetPost(data.PostID)
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
		return
	}

	err = utils.SetPostToCache(dbPost,c)
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, dbPost.Likes)
}
