package Oauth2

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type Client struct{

	Client oauth2.Config
	Resource map[string]string
	secret []byte
	store sessions.CookieStore
	state string
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
func (c *Client)Setup() error{
	c.secret = []byte("test")
	c.store = sessions.NewCookieStore(c.secret)


	return nil


}
func (c *Client)Session(name string) gin.HandlerFunc {
	return sessions.Sessions(name, c.store)
}
func (c *Client) LoginHandler(ctx *gin.Context) {
	c.state = randToken()
	session := sessions.Default(ctx)
	session.Set("state", c.state)
	session.Save()
	ctx.Writer.Write([]byte("<html><title>Golang Kakao</title> <body> <a href='" + c.GetLoginURL(c.state) + "'><button>Login with KaKao!</button> </a> </body></html>"))
}
func (c *Client)GetLoginURL(state string) string{
	return c.Client.AuthCodeURL(state)
}


