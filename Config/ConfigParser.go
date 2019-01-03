package Config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)



type Setting struct{

	Name string
	App Application
	Db Db
}

func (setting *Setting) getConfig() error{

	ymalFile, err := ioutil.ReadFile("./Config.yaml")
	if err != nil{
		log.Printf("yamlFile.Get err # %v ", err)
		return err
	}
	err = yaml.Unmarshal(ymalFile,&setting)
	if err != nil{
		log.Printf("Unmarshal: %v", err)
		return err
	}
	return nil

}

