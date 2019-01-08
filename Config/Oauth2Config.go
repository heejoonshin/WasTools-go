package Config

import (
	"github.com/heejoonshin/WasTools-go/Config/Oauth2/kakao"
	"github.com/heejoonshin/WasTools-go/Config/Oauth2/naver"
	"log"
)
type OauthMannager struct{
	Kakao kakao.KakaoOauth
	Naver naver.NaverOauth
}



func (oauthmannager *OauthMannager)ReadConfig(path string) error{

	err := oauthmannager.Kakao.ReadConfig(path)
	err = oauthmannager.Naver.ReadConfig(path)

	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	return nil

}
