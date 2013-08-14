package controllers

import (
	"github.com/astaxie/beego"
)

type AboutRouter struct {
	beego.Controller
}

func (this *AboutRouter) Get() {
	// Set const
	this.Data["App_Name"] = App_Name

	// Get User session
	u := this.GetSession("_User")
	if u == nil {
		this.Data["User"] = false
		this.TplNames = "about.html"
		return
	}
	User = u.(Session_User)
	this.Data["User"] = User

	this.TplNames = "about.html"
}
