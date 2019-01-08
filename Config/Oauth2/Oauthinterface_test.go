package Oauth2

import (
	"fmt"
	"testing"
)

func TestClient_ReadYaml(t *testing.T) {
	var oauthconf Oauthconf
	//oauthconf.ReadYaml("./kakao/kakao.yaml")
	fmt.Print(oauthconf)
}
