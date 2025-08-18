package managerControllers

import (
	"github.com/gin-gonic/gin"

	apiexception "CMS/app/apiException"
	managerservices "CMS/app/services/managerServices"
	"CMS/app/utils"
	"CMS/config/database"
)

type HandleData struct {
	UserID   int `json:"user_id" binding:"required"`
	ReportID int `json:"report_id" binding:"required"`
	Approval int `json:"approval" binding:"required"`
}

func ReportedPostHandling(c *gin.Context) {

	var data HandleData
	err := c.ShouldBindJSON(&data)
	if err != nil {

		apiexception.AbortWithException(c, apiexception.ParamError, err)
		return
	}
	flag, err := managerservices.ManagerJudge(data.UserID) //判断该用户是否为管理员
	if err != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, err)
		return
	}
	if !flag {
		apiexception.AbortWithException(c, apiexception.NotManagerError, err)
		return
	}

	// 开启事务
	tx := database.DB.Begin()
	if tx.Error != nil {
		apiexception.AbortWithException(c, apiexception.ServerError, tx.Error)
		return
	}

	defer func() {
		// 若发生 panic，回滚事务
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if data.Approval == 1 {

		PostID, err := managerservices.SeekPost(tx, data.ReportID)
		if err != nil {
			tx.Rollback() // 出错回滚
			apiexception.AbortWithException(c, apiexception.ServerError, err)
			return
		}

		err = managerservices.Delete(tx, PostID)
		if err != nil {
			tx.Rollback() // 出错回滚
			apiexception.AbortWithException(c, apiexception.ServerError, err)
			return
		}
	}

	if data.Approval == 1 || data.Approval == 2 {
		err := managerservices.Update(tx, data.ReportID, data.Approval)
		if err != nil {
			tx.Rollback() // 出错回滚
			apiexception.AbortWithException(c, apiexception.ServerError, err)
			return
		}
	} else {
		tx.Rollback()
		utils.JsonErrorResponse(c, 200509, "请输入正确审核结果")
		return
	}

	// 所有操作成功，提交事务
	_ = tx.Commit()

	utils.JsonSuccessResponse(c, nil)

}
