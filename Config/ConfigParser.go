package Config

import (
	"github.com/heejoonshin/WasTools-go/Config/Oauth2/custom"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)



type Setting struct{

	Name string
	App Application
	Db Db
	Oauth custom.CustomOauth
}

func (setting *Setting) GetConfig() error{


	ymalFile, err := ioutil.ReadFile("./Config/Config.yaml")
	if err != nil{
		log.Printf("yamlFile.Get err # %v ", err)
		return err
	}
	err = yaml.Unmarshal(ymalFile,&setting)
	if err != nil{
		log.Printf("Unmarshal: %v", err)
		return err
	}
	err = setting.Oauth.ReadConfig("./Config/Oauth2/oauth2.yaml")
	if err != nil{
		return err
	}
	return nil

}



