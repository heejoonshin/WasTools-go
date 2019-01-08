package naver

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/heejoonshin/WasTools-go/Config/Oauth2"
	"golang.org/x/oauth2"

	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)




type NaverOauth struct{
	Oauthconf Oauth2.Oauthconf `yaml:"naver"`
	secret []byte
	store sessions.CookieStore
	state string


}

func (k *NaverOauth)ReadConfig(path string) error{
	ymalFile, err := ioutil.ReadFile(path)
	if err != nil{
		log.Printf("yamlFile.Get err # %v ", err)
		return err
	}



	err = yaml.Unmarshal(ymalFile,&k)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	k.Setup()
	return nil

}

func (k *NaverOauth)Setup() error{
	k.secret = []byte("test")
	k.store = sessions.NewCookieStore(k.secret)
	return nil


}
func (c *NaverOauth)Session(name string) gin.HandlerFunc {
	return sessions.Sessions(name, c.store)
}
func (c *NaverOauth) LoginHandler(ctx *gin.Context) {
	c.state = Oauth2.RandToken()
	session := sessions.Default(ctx)
	session.Set("state", c.state)

	session.Save()
	ctx.Writer.Write([]byte("<html><title>Golang Kakao</title> <body> <a href='" + c.GetLoginURL(c.state) + "'><button>Login with KaKao!</button> </a> </body></html>"))
}


func (c *NaverOauth)GetLoginURL(state string) string{
	return c.Oauthconf.Client.AuthCodeURL(state)
}
func (c *NaverOauth)Auth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// Handle the exchange code to initiate a transport.
		session := sessions.Default(ctx)
		retrievedState := session.Get("state")
		retrievedState = strings.Replace(retrievedState.(string),"+"," ",-1)
		if retrievedState != ctx.Query("state") {
			ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
			return
		}

		tok, err := c.Oauthconf.Client.Exchange(oauth2.NoContext, ctx.Query("code"))
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(tok)
		//client := c.Oauthconf.Client.Client(oauth2.NoContext, tok)
		//email, err := client.Get()

		// save userinfo, which could be used in Handlers

	}

}


