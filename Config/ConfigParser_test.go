package Config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"testing"
)

var data = `
hostname: Easy!
c:
  c: 2
  d: [3, 4]
  z: 4
`

type T struct {
	Hostname string
	C struct {
		RenamedC int   `yaml:"c"`
		X int `yaml:"z"`
		D        []int `yaml:",flow"`
	}
}
func TestConfig(t *testing.T){
	var set Setting
	set.getConfig()

	fmt.Println(set)

}
func TestYaml(t *testing.T){
	tt := T{}

	err := yaml.Unmarshal([]byte(data), &tt)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("--- t:\n%v\n\n", t)
/*
	d, err := yaml.Marshal(&tt)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("--- t dump:\n%s\n\n", string(d))
	//fmt.Print(d)*/

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(tt)
	//fmt.Printf("--- m:\n%v\n\n", m)

/*
	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("--- m dump:\n%s\n\n", string(d))
*/
}
