package Config

import (
	"fmt"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)
type Oauth struct{

}

func Oauth2config() error{
	ymalFile, err := ioutil.ReadFile("./Oauth2/kakao.yaml")
	if err != nil{
		log.Printf("yamlFile.Get err # %v ", err)
		return err
	}

	var t oauth2.Config

	err = yaml.Unmarshal(ymalFile,&t)


	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Print(t)

	return nil
}

