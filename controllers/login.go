package controllers

import (
	"BTCsport/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strings"
)

type LoginRouter struct {
	beego.Controller
}

func (this *LoginRouter) Get() {
	// Set const
	this.Data["App_Name"] = App_Name
	// xsrf
	this.Data["xsrf"] = template.HTML(this.XsrfFormHtml())
	// set token in case twice submit
	Token := models.Token()
	this.SetSession("Token", Token)
	// set cookie not bot in case bots
	Cookies := models.RandString(20)
	this.SetSession("Cookie", Cookies)

	this.Data["Token"] = Token
	this.Data["Cookie"] = Cookies

	// Get referral
	Refer := this.Input().Get("username")
	if models.UserExist(Refer) {
		this.Data["Refer"] = Refer
	} else {
		models.Log(models.Log_Struct{"error", "Login:", errors.New("No such referer")})
	}

	// Get User Session
	var user Session_User
	u := this.GetSession("_User")
	if u != nil {
		user = u.(Session_User)
		this.Data["User"] = user
	} else {
		this.Data["User"] = false
	}

	// Show the login page or Error page
	var showLogin bool
	v := this.GetSession("ShowLogin")
	if v != nil {
		showLogin = v.(bool)
		this.DelSession("ShowLogin")
	} else {
		showLogin = true
	}
	this.Data["ShowLogin"] = showLogin

	// Errors
	if !showLogin {
		e := this.GetSession("Error")
		if e != nil {
			this.Data["Error"] = GetError(e)
			this.DelSession("Error")
		}
	}

	this.TplNames = "login.html"
}

func (this *LoginRouter) Post() {
	//fmt.Println(this.CheckXsrfCookie())

	// Get inputs and validate
	inputs := this.Input()

	// Check token in case twice submit
	var token string
	token_sess := this.GetSession("Token")
	if token_sess != nil {
		this.DelSession("Token")
		token = fmt.Sprintf("%d", token_sess.(int64))
	}
	if token != inputs.Get("token") {
		this.SetSession("Error", ERROR_TWICESUBMIT)
		models.Log(models.Log_Struct{"error", "Login:", errors.New("Submit twice")})
		this.fail()
		return
	}

	// Check cookie in case bots
	cookie_sess := this.GetSession("Cookie")
	if cookie_sess != nil {
		this.DelSession("Cookie")
		cookie := cookie_sess.(string)
		if cookie != this.Ctx.GetCookie("nobot") {
			this.SetSession("Error", ERROR_CAPTCHA)
			models.Log(models.Log_Struct{"error", "Login:", errors.New("No bot is allowed")})
			this.fail()
			return
		}
	}

	// Validate user inputs
	username := strings.TrimSpace(inputs.Get("username"))
	password := strings.TrimSpace(inputs.Get("password"))

	if models.ValidString(username) && models.ValidString(password) {
		// Check if user exist
		if !models.UserExist(username) {
			models.Log(models.Log_Struct{"info", "Login:", errors.New("User not exist")})
			this.SetSession("Error", ERROR_USERNOTEXIST)
			this.fail()
			return
		}
		// Check if password correct
		if !models.PassMatch(username, password) {
			models.Log(models.Log_Struct{"info", "Login:", errors.New(username + "Password incorrect")})
			this.SetSession("Error", ERROR_PASSINCORRECT)
			this.fail()
			return
		}

		// Update IP
		ip := strings.Split(this.Ctx.Request.RemoteAddr, ":")[0]
		models.UpdateIP(username, ip)

		// Get Uid
		id := models.UidByUsername(username)

		this.SetSession("_User", Session_User{Uid: id, Username: username, Ip: ip})
		this.succ()
		return
	}

	models.Log(models.Log_Struct{"info", "Login:", errors.New("Failed, invalid data")})
	this.SetSession("Error", ERROR_INVALIDINPUT)
	this.fail()
}

func (this *LoginRouter) fail() {
	this.SetSession("ShowLogin", false)
	this.Ctx.Redirect(302, "/login")
}

func (this *LoginRouter) succ() {
	this.Ctx.Redirect(302, "/")
}
