package controllers

import (
	"BTCsport/models"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
)

type WithdrawRouter struct {
	beego.Controller
}

func (this *WithdrawRouter) Get() {
	// Set const
	this.Data["App_Name"] = App_Name

	// xsrf
	this.Data["xsrf"] = template.HTML(this.XsrfFormHtml())

	// Get User session
	u := this.GetSession("_User")
	if u == nil {
		this.Ctx.Redirect(302, "/login")
		return
	}
	User := u.(Session_User)
	this.Data["User"] = User

	this.TplNames = "withdraw.html"
}

func (this *WithdrawRouter) Post() {
	// Get Data session: addr, authen
	d := this.GetSession("Data")
	var addr, authen, amount_str string
	if d == nil {
		this.SetSession("Error", ERROR_CAPTCHA)
		this.fail()
	}
	this.DelSession("Data")

	Data := d.([]string)
	addr = Data[0]
	authen = Data[1]
	amount_str = Data[2]

	// Validate user input authen
	if this.Input().Get("authen") != authen {
		this.SetSession("Error", ERROR_CAPTCHA)
		this.fail()
		return
	}

	// Get User session
	u := this.GetSession("_User")
	if u == nil {
		this.Ctx.Redirect(302, "/login")
		return
	}
	User := u.(Session_User)

	// Get User Address and Send BTC to User
	userAddr := models.AddressByUid(User.Uid)
	if userAddr == "" {
		this.SetSession("Error", ERROR_DB)
		this.fail()
		return
	}

	amount, _ := strconv.ParseFloat(amount_str, 64)
	_, err := models.SendBTC(userAddr, addr, "提现"+amount_str+"BTC", amount-Tx_Fee)
	if err != nil {
		this.SetSession("Error", ERROR_CANTSENDBTC)
		this.fail()
		return
	}

	// Update DB
	if !models.WithdrawRequest(User.Uid, amount, addr) {
		this.SetSession("Error", ERROR_DB)
		this.fail()
		return
	}

	this.succ()
}

func (this *WithdrawRouter) fail() {
	this.Ctx.Redirect(302, "/user/Withdraw")
}

func (this *WithdrawRouter) succ() {
	this.Ctx.Redirect(302, "/user/Withdraw")
}
