package managerservices

import (
	"CMS/app/models"
	"CMS/config/database"
)

func ManagerJudge(UserID int)(bool,error) {
	var data models.User
	result:=database.DB.Where("ID=?",UserID).First(&data)
	if result.Error!=nil{
		return false,result.Error
	}else if data.UserType==1{
		return false,nil
	}
	return true,nil
}

func ReportedPostShow()([]models.ReportedPost,error){
	reportedPosts:=[]models.ReportedPost{}
	result:=database.DB.Where("status=?",0).Find(&reportedPosts)
	if result.Error!=nil{
		return nil,result.Error
	}
	return reportedPosts,nil

}

func SeekPost(ReportID int)(int,error){
	var data models.ReportedPost
	result:=database.DB.Where("id=?",ReportID).First(&data)
	if result.Error!=nil{
		return 0,result.Error
	}
	return data.PostID,nil

}


func Update(ReportID int,Approval int)error{
	result:=database.DB.Model(&models.ReportedPost{}).Where("id=?",ReportID).Update("status",Approval)
	return result.Error
}
