package studentservices

import (
	"CMS/app/models"
	"CMS/config/database"
)

func ReleasePost(post *models.Post) error {
	result := database.DB.Create(post)
	return result.Error
}

func ShowPost(offset, pageSize int) ([]models.Post, error) {
	posts := []models.Post{}
	result := database.DB.Limit(pageSize).Offset(offset).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func Delete(id int) error {
	result := database.DB.Where("id=?", id).Delete(&models.Post{})
	return result.Error
}

func Update(data *models.Post) error {
	result := database.DB.Save(data)
	return result.Error
}

func AddReportedPost(ReportedPost *models.ReportedPost) error {
	result := database.DB.Create(ReportedPost)
	return result.Error

}

func GetPost(PostID int) (*models.Post, error) {
	var post models.Post
	result := database.DB.Where("ID=?", PostID).First(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil

}

func SeekUser(UserID int) (models.User, error) {
	var data models.User
	result := database.DB.Where("ID=?", UserID).First(&data)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return data, nil

}

func ShowReportedPost(UserID int) ([]models.ReportedPost, error) {
	posts := []models.ReportedPost{}
	result := database.DB.Where("user_id=?", UserID).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}
