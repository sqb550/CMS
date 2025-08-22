package studentControllers

import (
	apiexception "CMS/app/apiException"
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"
	
	"strconv"
	"strings"

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
	key := utils.GetPostCacheKey(uint(data.PostID))
	trimmed := strings.TrimPrefix(key, "post:")
    post_id, _:= strconv.ParseUint(trimmed, 10, 64)
	postID:=uint(post_id)
	post,ok,err:=utils.GetPostFromCache(postID)
	if err!=nil{
		apiexception.AbortWithException(c,apiexception.ServerError,err)
	}
	if !ok{
		dbPost,err:=studentservices.GetPost(data.PostID)
		if err!=nil{
			apiexception.AbortWithException(c,apiexception.ServerError,err)
			return
		}

		err = utils.SetPostToCache(dbPost)
		if err != nil {
			apiexception.AbortWithException(c,apiexception.ServerError,err)
			return
		}
	}
	post.Likes++
	err=utils.SetPostToCache(post)
	if err!=nil{
		apiexception.AbortWithException(c,apiexception.ServerError,err)
	}


	
	utils.JsonSuccessResponse(c,nil)
}

