package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorRouter struct {
	beego.Controller
}

func (this *ErrorRouter) Get() {
	this.TplNames = "404.html"
	this.Data["App_Name"] = App_Name
}
