package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/fzzy/sockjs-go/sockjs"
	"net/http"
	"strings"
	"time"
)

var Users *sockjs.SessionPool
var Conf sockjs.Config
var SockjsHandler http.Handler
var username string

func init() {
	Users = sockjs.NewSessionPool()
	Conf = sockjs.NewConfig()
	SockjsHandler = sockjs.NewHandler("/chat", ChatHandler, Conf)
}

func ChatHandler(s sockjs.Session) {
	Users.Add(s)
	defer Users.Remove(s)

	for {
		m := s.Receive()
		if m == nil {
			break
		}
		m = []byte(fmt.Sprintf("%s    %s:    %s", strings.Split(time.Now().Format(layout), " ")[1], username, m))
		Users.Broadcast(m)
		time.Sleep(5 * time.Second)
	}
}

type ChatRouter struct {
	beego.Controller
}

func (this *ChatRouter) Get() {
	// Set const
	this.Data["App_Name"] = App_Name

	// Get User session
	u := this.GetSession("_User")
	if u == nil {
		this.Data["User"] = false
		this.TplNames = "chat.html"
		return
	}
	User = u.(Session_User)
	this.Data["User"] = User
	username = User.Username

	this.TplNames = "chat.html"
}
