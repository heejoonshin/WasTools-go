package Oauth2

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Oauthconf struct{

	Client oauth2.Config
	Resource map[string]string

}


type Oauth2 interface {
	ReadConfig(path string) error
	Setup() error
	Session(name string) gin.HandlerFunc
	LoginHandler(ctx *gin.Context)
	GetLoginURL(state string) string
	Auth() gin.HandlerFunc

}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}




