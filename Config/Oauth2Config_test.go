package Config

import (
	"fmt"
	"golang.org/x/oauth2"
	"reflect"
	"testing"
)

func TestReadyaml(t *testing.T){
	var oauthtest OauthMannager
	err := oauthtest.ReadConfig("./Oauth2/oauth2.yaml")
	if err != nil{
		t.Fail()
	}
	fmt.Println(oauthtest)
}
func TestReflect(t *testing.T){
	var temp oauth2.Config

	types := reflect.ValueOf(&temp).Elem()
	for i := 0; i < types.NumField(); i++{
		Fieldname := types.Type().Field(i).Name
		FieldType := types.Type().Field(i).Type
		fmt.Println(Fieldname, FieldType)

	}
}