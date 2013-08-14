package controllers

import (
	"BTCsport/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strings"
)

type ResetRouter struct {
	beego.Controller
}

func (this *ResetRouter) Post() {
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
		models.Log(models.Log_Struct{"error", "Reset:", errors.New("Submit twice")})
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
			models.Log(models.Log_Struct{"error", "Reset:", errors.New("No bot is allowed")})
			this.fail()
			return
		}
	}

	// Get the email code
	var code string
	code_sess := this.GetSession("Authen")
	if code_sess != nil {
		this.DelSession("Authen")
		code = code_sess.(string)
	}
	// Get the username
	var username string
	user_sess := this.GetSession("Username")
	if user_sess != nil {
		this.DelSession("Username")
		username = user_sess.(string)
	}

	// Get user inputs
	authen := strings.TrimSpace(inputs.Get("authen"))
	password := strings.TrimSpace(inputs.Get("password"))
	re_password := strings.TrimSpace(inputs.Get("re-password"))
	fundpass := strings.TrimSpace(inputs.Get("fundpassword"))
	re_fundpass := strings.TrimSpace(inputs.Get("re-fundpassword"))

	// Validate user inputs
	if models.ValidString(password) && models.ValidString(re_password) && models.ValidString(fundpass) && models.ValidString(re_fundpass) && password == re_password && fundpass == re_fundpass && authen == code {
		// Check if code matches input
		if code != authen {
			models.Log(models.Log_Struct{"info", "Reset:", errors.New("Code not matches.")})
			this.SetSession("Error", ERROR_CODENOTMATCH)
			return
		}

		// Update DB
		if !models.UpdateUserPass(username, models.EncodePass(password), models.EncodePass(fundpass)) {
			models.Log(models.Log_Struct{"info", "Reset:", errors.New("Cant update password of user.")})
			this.SetSession("Error", ERROR_CANTUPDATE)
			return
		}

		this.succ()
		return
	}

	models.Log(models.Log_Struct{"info", "Reset:", errors.New("Failed, invalid data.")})
	this.SetSession("Error", ERROR_INVALIDINPUT)
	this.fail()
}

func (this *ResetRouter) Get() {
	// Set const
	this.Data["App_Name"] = App_Name
	// xsrf
	this.Data["xsrf"] = template.HTML(this.XsrfFormHtml())

	v := this.GetSession("Reset_Suc")
	if v != nil {
		this.DelSession("Reset_Suc")
		this.Data["Reset_Suc"] = v.(bool)
		if !v.(bool) {
			e := this.GetSession("Error")
			if e != nil {
				this.Data["Error"] = GetError(e)
				this.DelSession("Error")
			}
		}
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

	this.TplNames = "reset.html"
}

func (this *ResetRouter) succ() {
	this.SetSession("Reset_Suc", true)
	this.Ctx.Redirect(302, "/reset")
}

func (this *ResetRouter) fail() {
	this.SetSession("Reset_Suc", false)
	this.Ctx.Redirect(302, "/reset")
}
