package Config

import (
	"github.com/gin-gonic/gin"
	"github.com/heejoonshin/WasTools-go/Config/Oauth2"
	"github.com/zalando/gin-oauth2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

)
type OauthMannager struct{
	Kakao Oauth2.Client
}



func UidCheck(tc *ginoauth2.TokenContainer, ctx *gin.Context) bool {
	uid := tc.Scopes["uid"].(string)
	if uid != "sszuecs" {
		return false
	}
	ctx.Set("uid", uid)
	return true
}
func (oauthmannager *OauthMannager)ReadConfig(path string) error{
	ymalFile, err := ioutil.ReadFile(path)
	if err != nil{
		log.Printf("yamlFile.Get err # %v ", err)
		return err
	}
	err = yaml.Unmarshal(ymalFile,&oauthmannager)
	oauthmannager.Kakao.Setup()

	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	return nil

}
