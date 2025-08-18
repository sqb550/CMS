package managerservices

import (
	"CMS/app/models"
	"CMS/config/database"

	"gorm.io/gorm"
)

func ManagerJudge(UserID int) (bool, error) {
	var data models.User
	result := database.DB.Where("ID=?", UserID).First(&data)
	if result.Error != nil {
		return false, result.Error
	} else if data.UserType == 1 {
		return false, nil
	} else {
		return true, nil
	}
}

func ReportedPostShow() ([]models.ReportedPost, error) {
	reportedPosts := []models.ReportedPost{}
	result := database.DB.Where("status=?", 0).Find(&reportedPosts)
	if result.Error != nil {
		return nil, result.Error
	}
	return reportedPosts, nil

}

func SeekPost(tx *gorm.DB, ReportID int) (int, error) {
	var data models.ReportedPost
	result := tx.Where("id=?", ReportID).First(&data)
	if result.Error != nil {
		return 0, result.Error
	}
	return data.PostID, nil

}

func Update(tx *gorm.DB, ReportID int, Approval int) error {
	result := tx.Model(&models.ReportedPost{}).Where("id=?", ReportID).Update("status", Approval)
	return result.Error
}

func Delete(tx *gorm.DB, PostID int) error {
	result := tx.Where("id=?", PostID).Delete(&models.Post{})
	return result.Error
}
