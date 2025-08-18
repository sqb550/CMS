package router

import (
	"github.com/gin-gonic/gin"

	"CMS/app/controllers/managerControllers"
	"CMS/app/controllers/studentControllers"
	usercontrollers "CMS/app/controllers/userControllers"
	"CMS/app/midwares"
)

func Init(r *gin.Engine) {
	const pre = "/api"
	api := r.Group(pre)
	{
		// User 路由组
		user := api.Group("/user")
		{
			user.POST("/login", usercontrollers.Login)
			user.POST("/reg", usercontrollers.Register)
		}

		// Student 路由组
		student := api.Group("/student")
		student.Use(midwares.AuthMiddleware())
		{
			student.POST("/post", studentControllers.Release)
			student.GET("/post", studentControllers.Show)
			student.DELETE("/post", studentControllers.Delete)
			student.PUT("/post", studentControllers.Update)
			student.POST("/report-post", studentControllers.Report)
			student.GET("/report-post", studentControllers.ShowReportedPost)
		}

		// Admin 路由组
		admin := api.Group("/admin")
		admin.Use(midwares.AuthMiddleware())
		{
			admin.GET("/report", managerControllers.ShowReportedPosts)
			admin.POST("/report", managerControllers.ReportedPostHandling)
		}
	}
}
