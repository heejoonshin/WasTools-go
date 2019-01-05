package main

import (
	"github.com/gin-gonic/gin"
	"github.com/szuecs/gin-glog"
	"github.com/zalando/gin-oauth2"
	"time"
)

func main() {
	router := gin.New()
	router.Use(ginglog.Logger(3 * time.Second))
	router.Use(gin.Recovery())

	ginoauth2.VarianceTimer = 300 * time.Millisecond

	public := router.Group("/api")
	public.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{"message": "Hello to public world"})
	})



}
