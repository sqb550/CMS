package managercontrollers

import (
	managerservices "CMS/app/services/managerServices"
	studentservices "CMS/app/services/studentServices"
	"CMS/app/utils"

	"github.com/gin-gonic/gin"
)

type HandleData struct {
	UserID   int `json:"user_id" binding:"required"`
	ReportID   int `json:"report_id" binding:"required"`
	Approval int `json:"approval" binding:"required"`
}

func ReportedPostHandling(c *gin.Context){
	var data HandleData
	err:=c.ShouldBindJSON(&data)
	if err!=nil{
		utils.JsonErrorResponse(c,200501,"参数错误")
		return
	}
	flag,err:=managerservices.ManagerJudge(data.UserID)//判断该用户是否为管理员
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	if !flag{
		utils.JsonErrorResponse(c,200508,"非管理员不具备审核权限")
		return
	}
	if data.Approval==1{
		PostID,err:=managerservices.SeekPost(data.ReportID)
		if err!=nil{
			utils.JsonInternalServerErrorResponse(c)
			return
		}
		err=studentservices.Delete(PostID)
		if err!=nil{
			utils.JsonInternalServerErrorResponse(c)//判断处理结果是否为同意，同意则删除在post中的该帖子
			return
		}

	}
	if data.Approval==1 || data.Approval==2{
		err=managerservices.Update(data.ReportID,data.Approval)//判断处理结果是否为1或2
		if err!=nil{
			utils.JsonInternalServerErrorResponse(c)
			return
		}

	}else{
		utils.JsonErrorResponse(c,200509,"请输入正确审核结果")
	}
	
	utils.JsonSuccessResponse(c,nil)

	
}