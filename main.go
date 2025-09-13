package main

import (
	"log"

	"CMS/app/midwares"
	"CMS/app/utils"
	"CMS/config/database"
	"CMS/config/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	database.Init()
	utils.InitRedis()

	r := gin.Default()
	store := memstore.NewStore([]byte("secretkey"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(midwares.ErrHandler())

	c:=cron.New()
	_,err:=c.AddFunc("**/5****",utils.SyncCacheToDB)
	c.Start()
	defer c.Stop()

	router.Init(r)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Server start error:", err)
	}
}
