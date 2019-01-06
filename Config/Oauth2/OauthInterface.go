package Oauth2

import (
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Client struct{

	Client oauth2.Config
	Resource map[string]string
}

func (c *Client)ReadYaml(path string) error{

	ymalFile, err := ioutil.ReadFile(path)
	if err != nil{
		log.Printf("yamlFile.Get err # %v ", err)
		return err
	}
	err = yaml.Unmarshal(ymalFile,&c)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}

	return nil


}


