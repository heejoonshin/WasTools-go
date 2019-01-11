package Oauth2

import (
	"github.com/gin-gonic/gin"
)




type Oauth2 interface {

	Setup() error
	Session(name string) gin.HandlerFunc
	LoginHandler(ctx *gin.Context)
	GetLoginURL(state string) string
	Auth() gin.HandlerFunc

}





