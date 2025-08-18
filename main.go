package main

import (
	"log"

	"CMS/app/midwares"
	"CMS/config/database"
	"CMS/config/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()
	store := memstore.NewStore([]byte("secretkey"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(midwares.ErrHandler())

	router.Init(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Server start error:", err)
	}
}
