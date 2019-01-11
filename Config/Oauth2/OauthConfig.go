package Oauth2

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"strings"
)
type Oauthconf struct{


	ClientName string
	Client oauth2.Config
	Resource map[string]string
	secret []byte
	store sessions.CookieStore
	state string

}
func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}


func (o *Oauthconf)Setup() error{
	o.secret = []byte("test")
	o.store = sessions.NewCookieStore(o.secret)
	return nil
}
func (o *Oauthconf)Session(name string) gin.HandlerFunc{
	return sessions.Sessions(name,o.store)
}
func (o *Oauthconf)LoginHandler(ctx *gin.Context){
	o.state = RandToken()
	session := sessions.Default(ctx)
	session.Set("state", o.state)
	session.Save()
	ctx.Writer.Write([]byte("<html><title>Golang "+o.ClientName+"</title> <body> <a href='" + o.GetLoginURL(o.state) + "'><button>Login with "+o.ClientName+"</button> </a> </body></html>"))
}

func (o *Oauthconf)GetLoginURL(state string) string{
	return o.Client.AuthCodeURL(state)
}
func (o *Oauthconf)Auth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// Handle the exchange code to initiate a transport.
		session := sessions.Default(ctx)
		retrievedState := session.Get("state")
		retrievedState = strings.Replace(retrievedState.(string),"+"," ",-1)
		if retrievedState != ctx.Query("state") {
			ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
			return
		}

		tok, err := o.Client.Exchange(oauth2.NoContext, ctx.Query("code"))
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(tok)

		client := o.Client.Client(oauth2.NoContext, tok)
		userinfo, err := client.Get(o.Resource["userinfo"])
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		defer userinfo.Body.Close()


		data, err := ioutil.ReadAll(userinfo.Body)

		if err != nil {
			glog.Errorf("[Gin-OAuth] Could not read Body: %s", err)
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}


		var user map[string]interface{}
		err = json.Unmarshal(data,&user)

		if err != nil {
			glog.Errorf("[Gin-OAuth] Unmarshal userinfo failed: %s", err)
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		fmt.Println(user)
		ctx.Set("user", user)




		// save userinfo, which could be used in Handlers

	}

}
