package controllers

import (
	"BTCsport/models"
	// "fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
	"time"
)

type GameRouter struct {
	beego.Controller
}

func (this *GameRouter) Get() {

	// Set App_Name
	this.Data["App_Name"] = App_Name

	// XSRF
	this.Data["xsrf"] = template.HTML(this.XsrfFormHtml())

	// Get User Session
	var user Session_User
	u := this.GetSession("_User")
	if u != nil {
		user = u.(Session_User)
		this.Data["User"] = user
	} else {
		this.Data["User"] = false
	}

	// Get id
	id, _ := strconv.Atoi(this.Ctx.Params[":id"])

	// Get game
	game := models.GameById(id)

	// Check if users can bet
	timeSub := TimeSub(game.Timestarted).Minutes()
	if game.Isover == 1 || timeSub < 15 {
		this.Data["CanBet"] = false
	} else {
		this.Data["CanBet"] = true
	}

	// Set Error alert (From Post)
	this.Data["PostResult"] = false
	v := this.GetSession("BetSuccess")
	if v != nil {
		this.Data["PostResult"] = true
		this.DelSession("BetSuccess")
		this.Data["BetSuccess"] = v.(bool)
		if !v.(bool) {
			e := this.GetSession("Error")
			this.DelSession("Error")
			if e != nil {
				this.Data["Error"] = GetError(e)
			}
		}
	}

	//  Set game info
	this.Data["Game"] = game

	this.TplNames = "game.html"
}

func (this *GameRouter) Post() {

	// Get _User session --> uid
	user_sess := this.GetSession("_User")
	if user_sess == nil {
		this.Ctx.Redirect(302, "/login")
		return
	}
	user := user_sess.(Session_User)

	// Get uid
	uid := user.Uid

	// Get game id
	id, _ := strconv.Atoi(this.Ctx.Params[":id"])

	// Post type
	Type := this.Input().Get("type")

	// Deal with post
	this.postData(Type, id, uid)

}

func (this *GameRouter) postData(Type string, id, uid int) {
	// Validate Bet amount
	amount_str := this.Input().Get("amount")
	if !models.ValidBetamount(amount_str) {
		this.SetSession("Error", ERROR_INVALIDINPUT)
		this.fail(id)
		return
	}

	// Validate if balance is enough
	amount, _ := strconv.ParseFloat(amount_str, 64)
	if !models.UserBalanceEnough(uid, amount) {
		this.SetSession("Error", ERROR_BALANCENOTENOUGH)
		this.fail(id)
		return
	}

	// Submit Post
	switch Type {
	case "AorB":
		// Validate radio input
		sliceAorB := []interface{}{"A_Win", "B_Win"}
		AorB := this.Input().Get("AorB")
		if !In_slice(AorB, sliceAorB) {
			this.SetSession("Error", ERROR_INVALIDINPUT)
			this.fail(id)
			return
		}

		// Send BTC to betaddress
		Type := " 主队赢"
		if AorB != "A_Win" {
			Type = " 客队赢"
		}
		user := models.UserByUid(uid)
		note := time.Now().Format(layout) + " 赛事ID:" + strconv.Itoa(id) + Type
		txhash, err := models.SendBTC(user.Btcaddress, BetAddr, note, amount)
		if err != nil {
			this.SetSession("Error", ERROR_CANTSENDBTC)
			this.fail(id)
			return
		}

		// Change database
		if !models.PlaceBet(id, uid, AorB, txhash, amount) {
			this.SetSession("Error", ERROR_DB)
			this.fail(id)
			return
		}

	case "OddEven":
		sliceOddEven := []interface{}{"Odd", "Even"}
		OddEven := this.Input().Get("OddEven")

		if !In_slice(OddEven, sliceOddEven) {
			this.SetSession("Error", ERROR_INVALIDINPUT)
			this.fail(id)
			return
		}

		// Send BTC to betaddress
		Type := " 单数总比分"
		if OddEven != "Odd" {
			Type = " 双数总比分"
		}
		user := models.UserByUid(uid)
		note := time.Now().Format(layout) + " 赛事ID:" + strconv.Itoa(id) + Type
		txhash, err := models.SendBTC(user.Btcaddress, BetAddr, note, amount)
		if err != nil {
			this.SetSession("Error", ERROR_CANTSENDBTC)
			this.fail(id)
			return
		}

		// Change database
		if !models.PlaceBet(id, uid, OddEven, txhash, amount) {
			this.SetSession("Error", ERROR_DB)
			this.fail(id)
			return
		}

	case "LargeSmall":
		sliceLargeSmall := []interface{}{"Large", "Small"}
		LargeSmall := this.Input().Get("LargeSmall")

		if !In_slice(LargeSmall, sliceLargeSmall) {
			this.SetSession("Error", ERROR_INVALIDINPUT)
			this.fail(id)
			return
		}

		// Send BTC to betaddress
		Type := " 大分"
		if LargeSmall != "Large" {
			Type = " 小分"
		}
		user := models.UserByUid(uid)
		note := time.Now().Format(layout) + " 赛事ID:" + strconv.Itoa(id) + Type
		txhash, err := models.SendBTC(user.Btcaddress, BetAddr, note, amount)
		if err != nil {
			this.SetSession("Error", ERROR_CANTSENDBTC)
			this.fail(id)
			return
		}

		// Change database
		if !models.PlaceBet(id, uid, LargeSmall, txhash, amount) {
			this.SetSession("Error", ERROR_DB)
			this.fail(id)
			return
		}

	}

	this.succ(id)
}

func (this *GameRouter) fail(id int) {
	this.SetSession("BetSuccess", false)
	this.Ctx.Redirect(302, "/game/"+strconv.Itoa(id))
}

func (this *GameRouter) succ(id int) {
	this.SetSession("BetSuccess", true)
	this.Ctx.Redirect(302, "/game/"+strconv.Itoa(id))
}
