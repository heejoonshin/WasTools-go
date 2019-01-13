package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/heejoonshin/WasTools-go/Config"
	"github.com/szuecs/gin-glog"
	"github.com/zalando/gin-oauth2"
	"time"
)

func main() {
	var property Config.Setting
	property.GetConfig()

	router := gin.New()
	router.Use(ginglog.Logger(3 * time.Second))
	router.Use(gin.Recovery())

	ginoauth2.VarianceTimer = 300 * time.Millisecond

	public := router.Group("/api")
	public.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello to public world"})
	})
	kakaoOauth, _ := property.Oauth.GetOauthProvider("kakao")
	kakaoOauth.Setup()

	router.Use(kakaoOauth.Session("goquestsession"))
	router.GET("/login", kakaoOauth.LoginHandler)

	private := router.Group("/auth")

	private.Use(kakaoOauth.Auth())
	//private.Use(property.Oauth.Naver.Auth())
	private.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from private"})
	})
	glog.Info("bootstrapped application")

	router.Run(":8080")

}
