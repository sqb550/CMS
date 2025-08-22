package studentControllers

import (
	apiexception "CMS/app/apiException"
	"CMS/app/models"
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReleaseData struct {
	Content string `json:"content" binding:"required"`
}
type PostList struct {
	PostList []models.Post `json:"post_list"`
}

type PageQuery struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
}

type DeleteData struct {
	PostID int `form:"post_id" json:"id" binding:"required"`
}

type UpdateData struct {
	ID      uint   `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func Release(c *gin.Context) {
	UserID, exists := c.Get("user_id")
	if !exists {
		apiexception.AbortWithException(c, apiexception.USerIDError, nil)
	}
	UserIDInt, _ := UserID.(int)

	var data ReleaseData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ParamError, err)
		return
	}
	//将帖子信息添加到数据表post
	err = studentservices.ReleasePost(&models.Post{
		Content: data.Content,
		UserID:  UserIDInt,
	})
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}

func Show(c *gin.Context) {
	var data PageQuery
	err := c.ShouldBindQuery(&data)
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
	}
	offset := (data.Page - 1) * data.PageSize                      //计算偏移量
	result, err := studentservices.ShowPost(offset, data.PageSize) //result为post中的结构体数组
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
	}
	var results PostList
	results.PostList = result
	utils.JsonSuccessResponse(c, results)
}

func Delete(c *gin.Context) {

	var data DeleteData
	err := c.ShouldBindQuery(&data) //绑定请求的参数到结构体中
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ParamError, err)
		return
	}

	err = studentservices.Delete(data.PostID) //删除某一条帖子
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			apiexception.AbortWithException(c, apiexception.PostNotFound, err)
		} else {
			apiexception.AbortWithException(c, apiexception.ServerError, err)
		}
		return

	}
	utils.JsonSuccessResponse(c, nil)

}

func Update(c *gin.Context) {
	UserID, exists := c.Get("user_id")
	if !exists {
		apiexception.AbortWithException(c, apiexception.USerIDError, nil)
	}
	UserIDInt, _ := UserID.(int)
	var data UpdateData
	err := c.ShouldBindJSON(&data) //绑定
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ParamError, err)
		return
	}

	err = studentservices.Update(&models.Post{
		ID:      data.ID,
		UserID:  UserIDInt,
		Content: data.Content,
	}) //更新数据
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200510, "该帖子不存在")
		} else {
			apiexception.AbortWithException(c, apiexception.ServerError, err)
		}
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
