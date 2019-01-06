package Config

import (
	"github.com/heejoonshin/WasTools-go/Config/Oauth2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)
type OauthMannager struct{
	Kakao Oauth2.Client
}
func (oauthmannager *OauthMannager)ReadConfig(path string) error{
	ymalFile, err := ioutil.ReadFile(path)
	if err != nil{
		log.Printf("yamlFile.Get err # %v ", err)
		return err
	}
	err = yaml.Unmarshal(ymalFile,&oauthmannager)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	return nil

}
