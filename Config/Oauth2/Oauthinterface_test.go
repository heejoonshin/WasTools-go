package Oauth2

import (
	"fmt"
	"testing"
)

func TestClient_ReadYaml(t *testing.T) {
	var oauthconf Client
	//oauthconf.ReadYaml("./kakao/kakao.yaml")
	fmt.Print(oauthconf)
}
