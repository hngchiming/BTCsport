package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	// Set const
	this.Data["App_Name"] = App_Name
	this.Data["BetAddr"] = BetAddr
	this.Data["ProfitAddr"] = ProfitAddr

	// xsrf
	this.Data["xsrf"] = template.HTML(this.XsrfFormHtml())

	// Check if logout
	if this.Input().Get("logout") == "true" {
		this.DelSession("_User")
		this.Ctx.Redirect(302, "/")
		return
	}

	// Get User session
	var user Session_User
	u := this.GetSession("_User")
	if u != nil {
		user = u.(Session_User)
		this.Data["User"] = user
	} else {
		this.Data["User"] = false
	}

	this.TplNames = "index.tpl"
}
