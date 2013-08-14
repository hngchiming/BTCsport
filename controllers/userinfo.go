package controllers

import (
	"BTCsport/models"
	"github.com/astaxie/beego"
	"html/template"
	"math"
	"strconv"
	"strings"
)

type UserRouter struct {
	beego.Controller
}

var (
	User      Session_User
	User_Info models.User
	Deposit   models.Deposit
	Withdraw  models.Withdraw
)

func (this *UserRouter) Get() {

	// Set const
	this.Data["App_Name"] = App_Name

	// xsrf
	this.Data["xsrf"] = template.HTML(this.XsrfFormHtml())

	// Get User session
	u := this.GetSession("_User")
	if u == nil {
		this.Data["User"] = false
		this.TplNames = "user_Deposit.tpl"
		return
	}
	User = u.(Session_User)
	this.Data["User"] = User

	// Request URL
	reqUrl := this.Ctx.Request.URL.String()
	sec := reqUrl[strings.LastIndex(reqUrl, "/")+1:]
	if qm := strings.Index(sec, "?"); qm > -1 {
		sec = sec[:qm]
	}

	if len(sec) == 0 || sec == "user" {
		sec = "Deposit"
		this.Data[sec] = true
	} else {
		this.Data[sec] = true
	}

	this.setData(sec)
}

func (this *UserRouter) setData(sec string) {
	switch sec {
	case "Deposit":
		// User Addr
		this.Data["Addr"] = models.AddressByUid(User.Uid)

		// User All deposit
		userDeposit := models.DepositByUid(User.Uid)
		this.Data["UserDeposit"] = userDeposit

		// All time deposit
		var allTimeDeposit float64
		for _, v := range userDeposit {
			allTimeDeposit += v.Amount
		}
		this.Data["AllTimeDeposit"] = allTimeDeposit

		// Current Balance
		this.Data["Balance"] = models.UserByUid(User.Uid).Balance

	case "Withdraw":
		// Withdraw error from POST
		e := this.GetSession("Error")
		if e != nil {
			this.DelSession("Error")
			this.Data["Error"] = GetError(e)
		}

		// User All withdraw
		userWithdraw := models.WithdrawByUid(User.Uid)
		this.Data["UserWithdraw"] = userWithdraw

		//xsrf
		this.Data["xsrf"] = template.HTML(this.XsrfFormHtml())
	case "Refer":
		// Username -> Refer Address
		this.Data["Username"] = User.Username

		// User Bet amount and profit of this month
		betAmount, betProfit := models.UserBetAmountProfit(User.Uid)
		this.Data["betAmount"] = betAmount
		this.Data["betProfit"] = betProfit

		// Find all referee
		userRefer := models.UserByReferral(User.Username)
		this.Data["Count"] = len(userRefer)

		// Calculate all referees' bet info
		type temp struct {
			Uid       int
			Betamount float64
			Profit    float64
		}
		var userData []temp
		for _, v := range userRefer {
			betamount, profit := models.UserBetAmountProfit(v.Id)
			userData = append(userData, temp{Uid: v.Id, Betamount: betamount, Profit: profit})
		}
		this.Data["UserData"] = userData

		// Calculate ReferFee
		var referFee float64
		for _, v := range userData {
			if v.Profit < 0 {
				fee := math.Floor(0.00188*float64(len(userRefer))*100000000) / 10000000
				referFee -= v.Profit * fee
			}
		}
		this.Data["ReferFee"] = referFee

	case "Allbet":
		bets := models.UserBets(User.Uid)
		this.Data["BetCount"] = len(bets)
		this.Data["UserAllBets"] = bets

	}
	this.TplNames = "user_" + sec + ".tpl"
}

func (this *UserRouter) Post() {

	// Get _User session --> uid
	user_sess := this.GetSession("_User")
	if user_sess == nil {
		this.Ctx.Redirect(302, "/login")
		return
	}
	user := user_sess.(Session_User)

	// Get username
	username := user.Username

	// Get user inputs
	inputs := this.Input()
	email := inputs.Get("email")
	addr := inputs.Get("address")
	amount := inputs.Get("amount")
	fdps := inputs.Get("fundpassword")
	authen := models.RandString(15)

	// Validate inputs
	if models.ValidString(fdps) && models.ValidEmail(email) && models.ValidBetamount(amount) {

		// Validate fundpass
		if !models.FundPassMatch(username, fdps) {
			this.SetSession("Error", ERROR_PASSINCORRECT)
			this.fail()
			return
		}

		// Validate Email
		if !models.EmailMatch(username, email) {
			this.SetSession("Error", ERROR_EMAILNOTMATCH)
			this.fail()
			return
		}

		// Check if balance enough
		amount_float64, _ := strconv.ParseFloat(amount, 64)
		if !models.UserBalanceEnough(user.Uid, amount_float64) {
			this.SetSession("Error", ERROR_BALANCENOTENOUGH)
			this.fail()
			return
		}

		// Send Email Code to User
		if !models.SendEmail(email, "提现申请", "申请提现"+amount+"BTC到以下地址"+addr+"\n请复制右边的Code，以完成提现操作", authen) {
			this.SetSession("Error", ERROR_EMAILNOTSENT)
			this.fail()
			return
		}

		this.SetSession("Data", []string{addr, authen, amount})
		this.succ()
		return
	}

	this.SetSession("Error", ERROR_INVALIDINPUT)
	this.fail()
	return
}

func (this *UserRouter) fail() {
	this.Ctx.Redirect(302, "/user/Withdraw")
}

func (this *UserRouter) succ() {
	this.Ctx.Redirect(302, "/withdraw")
}
