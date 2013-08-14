package controllers

import (
	"BTCsport/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strings"
)

type ForgetRouter struct {
	beego.Controller
}

func (this *ForgetRouter) Post() {
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
		models.Log(models.Log_Struct{"error", "Forget:", errors.New("Submit twice")})
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
			models.Log(models.Log_Struct{"error", "Forget:", errors.New("No bot is allowed")})
			this.fail()
			return
		}
	}

	// Validate inputs
	username := strings.TrimSpace(inputs.Get("username"))
	dateofbirth := strings.TrimSpace(inputs.Get("birth"))
	email := strings.TrimSpace(inputs.Get("email"))

	if models.ValidString(username) && models.ValidEmail(email) && models.ValidBirth(dateofbirth) {
		// Check if user exist
		if !models.UserExist(username) {
			models.Log(models.Log_Struct{"info", "Forget:", errors.New("User not exist.")})
			this.SetSession("Error", ERROR_USERNOTEXIST)
			this.fail()
			return
		}
		// Check if birth matches
		if !models.BirthMatch(username, dateofbirth) {
			models.Log(models.Log_Struct{"info", "Forget:", errors.New("Birth not match.")})
			this.SetSession("Error", ERROR_BIRTHNOTMATCH)
			this.fail()
			return
		}
		// Check if email mathces
		if !models.EmailMatch(username, email) {
			models.Log(models.Log_Struct{"info", "Forget:", errors.New("Email not match.")})
			this.SetSession("Error", ERROR_EMAILNOTMATCH)
			this.fail()
			return
		}

		// Send Email to authenticate
		authen := models.RandString(8)
		if !models.SendEmail(email, "重设密码", username+":  请复制验证码，以完成重设密码操作---->", authen) {
			models.Log(models.Log_Struct{"warn", "Forget:", errors.New("Cant send email to authen password reset.")})
			this.SetSession("Error", ERROR_EMAILNOTSENT)
			this.fail()
			return
		}

		this.SetSession("Username", username)
		this.SetSession("Authen", authen)
		this.succ()
		return
	}

	models.Log(models.Log_Struct{"info", "Forget:", errors.New("Failed, invalid data.")})
	this.SetSession("Error", ERROR_INVALIDINPUT)
	this.fail()
}

func (this *ForgetRouter) Get() {
	// Set const
	this.Data["App_Name"] = App_Name

	// xsrf
	this.Data["xsrf"] = template.HTML(this.XsrfFormHtml())

	Token := models.Token()
	this.SetSession("Token", Token)
	this.Data["Token"] = Token

	v := this.GetSession("Forget_Suc")
	if v != nil {
		this.Data["Forget_Suc"] = v.(bool)
		if !v.(bool) {
			e := this.GetSession("Error")
			if e != nil {
				this.Data["Error"] = GetError(e)
				this.DelSession("Error")
			}
		}
		this.DelSession("Forget_Suc")
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

	this.TplNames = "forget.html"
}

func (this *ForgetRouter) fail() {
	this.SetSession("Forget_Suc", false)
	this.Ctx.Redirect(302, "/forget")
}

func (this *ForgetRouter) succ() {
	this.SetSession("Forget_Suc", true)
	this.Ctx.Redirect(302, "/forget")
}
