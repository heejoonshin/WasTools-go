package custom

import (
	"fmt"
	"github.com/heejoonshin/WasTools-go/Config/Oauth2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type CustomOauth struct{
	oauth map[string]*Oauth2.Oauthconf
}
func NewCustom() *CustomOauth{
	ret := &CustomOauth{}
	ret.oauth = make(map[string]*Oauth2.Oauthconf)
	return ret
}

func (o *CustomOauth)ReadConfig(path string)error{

	yamlFile, err:= ioutil.ReadFile(path)
	if err != nil{
		log.Printf("yamlFile.Get err # %v ", err)
		return err
	}
	configMap := make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile,&configMap)
	if err != nil{
		return err
	}
	err = o.customOauthParser(configMap)
	if err != nil{
		return err
	}
	return nil

}
func (o *CustomOauth)customOauthParser(config map[string]interface{}) (err error){
	//fmt.Println(config)
	for k,v := range config{
		fmt.Println(v)
		var b []byte
		b,err = yaml.Marshal(v)
		if err != nil{
			return nil
		}
		o.oauth[k] = &Oauth2.Oauthconf{}
		o.oauth[k].ClientName = k
		err = yaml.Unmarshal(b,o.oauth[k])
		if err != nil{
			return err
		}
	}
	return nil
}
func (o *CustomOauth)GetOauthProvider(provider string) (oauth *Oauth2.Oauthconf, err error){
	if value,ok := o.oauth[provider]; ok {
		return value,nil

	}else{
		return nil,fmt.Errorf("Not Found Provider")

	}

}