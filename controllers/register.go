package controllers

import (
	"BTCsport/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strings"
)

type RegisterRouter struct {
	beego.Controller
}

func (this *RegisterRouter) Get() {
	// Set const
	this.Data["App_Name"] = App_Name

	// xsrf
	this.Data["xsrf"] = template.HTML(this.XsrfFormHtml())

	v := this.GetSession("Reg_Success")
	if v != nil {
		this.Data["Reg_Success"] = v.(bool)
		if !v.(bool) {
			e := this.GetSession("Error")
			if e != nil {
				this.Data["Error"] = GetError(e)
				this.DelSession("Error")
			}
		}
		this.DelSession("Reg_Success")
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

	this.TplNames = "register.html"
}

func (this *RegisterRouter) Post() {
	// Get user inputs
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
		models.Log(models.Log_Struct{"error", "Register:", errors.New("Submit twice")})
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
			models.Log(models.Log_Struct{"error", "Register:", errors.New("No bot is allowed")})
			this.fail()
			return
		}
	}

	// Validate user inputs
	username := strings.TrimSpace(inputs.Get("username"))
	password := strings.TrimSpace(inputs.Get("password"))
	re_password := strings.TrimSpace(inputs.Get("re-password"))
	fundpass := strings.TrimSpace(inputs.Get("fundpassword"))
	re_fundpass := strings.TrimSpace(inputs.Get("re-fundpassword"))
	email := strings.TrimSpace(inputs.Get("email"))
	dateofbirth := strings.TrimSpace(inputs.Get("birth"))

	// Check referral
	refer := strings.TrimSpace(inputs.Get("refer"))
	if refer != "" && models.ValidString(refer) {
		if !models.UserExist(refer) {
			models.Log(models.Log_Struct{"info", "Register:", errors.New("Referral user not exist.")})
			this.SetSession("Error", ERROR_REFERNOTEXIST)
			this.fail()
			return
		}
	}

	// Validate user inputs, set sessions and redirect
	if models.ValidString(username) && models.ValidString(password) && models.ValidString(re_password) && models.ValidString(fundpass) && models.ValidString(re_fundpass) && password == re_password && fundpass == re_fundpass && models.ValidEmail(email) && models.ValidBirth(dateofbirth) {
		// Check if user exist
		if models.UserExist(username) {
			models.Log(models.Log_Struct{"info", "Register:", errors.New("User already exist.")})
			this.SetSession("Error", ERROR_USEREXIST)
			this.fail()
			return
		}

		// Check if email exist
		if models.EmailExist(email) {
			models.Log(models.Log_Struct{"info", "Register:", errors.New("Email already exist.")})
			this.SetSession("Error", ERROR_EMAILEXIST)
			this.fail()
			return
		}

		// Generate new address for new user
		address, err := models.NewAddress(username)
		if err != nil {
			models.Log(models.Log_Struct{"info", "Register:", err})
			this.fail()
			this.SetSession("Error", ERROR_ADDRESS)
			return
		}

		// Insert new user to DB
		ok := models.NewUser(models.User{Username: username, Password: models.EncodePass(password), Fundpassword: models.EncodePass(fundpass), Email: email, Btcaddress: address, Birth: dateofbirth, Referral: refer})
		if !ok {
			models.Log(models.Log_Struct{"info", "Register:", errors.New("Unable to insert user, need to delete from wallet.")})
			err = models.Archive(address)
			if err != nil {
				models.Log(models.Log_Struct{"warn", "Register:", errors.New("Unable to archive.")})
			} else {
				models.Log(models.Log_Struct{"info", "Register:", errors.New("Succeed archiving address.")})
			}
			this.fail()
			this.SetSession("Error", ERROR_DB)
			return
		}

		this.succ()
		return
	}

	models.Log(models.Log_Struct{"info", "Register:", errors.New("Failed, invalid data")})
	this.SetSession("Error", ERROR_INVALIDINPUT)
	this.fail()
}

func (this *RegisterRouter) fail() {
	this.SetSession("Reg_Success", false)
	this.Ctx.Redirect(302, "/register")
}

func (this *RegisterRouter) succ() {
	this.SetSession("Reg_Success", true)
	this.DelSession("_User")
	this.Ctx.Redirect(302, "/register")
}
