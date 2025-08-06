package router

import (
	"CMS/app/controllers/managercontrollers"
	"CMS/app/controllers/studentcontrollers"
	usercontrollers "CMS/app/controllers/userControllers"

	"github.com/gin-gonic/gin"
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
        {
            student.POST("/post", studentcontrollers.Release)
            student.GET("/post", studentcontrollers.Show)
            student.DELETE("/post", studentcontrollers.Delete)
            student.PUT("/post", studentcontrollers.Update)
            student.POST("/report-post", studentcontrollers.Report)
            student.GET("/report-post", studentcontrollers.ShowReportedPost)
        }

        // Admin 路由组
        admin := api.Group("/admin")
        {
            admin.GET("/report", managercontrollers.ShowReportedPosts)
            admin.POST("/report", managercontrollers.ReportedPostHandling)
        }
    }
}
