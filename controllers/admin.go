package controllers

import (
	"BTCsport/models"
	"errors"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
	"strings"
	"time"
)

type AdminRouter struct {
	beego.Controller
}

func (this *AdminRouter) Get() {
	// Set AppName
	this.Data["App_Name"] = App_Name

	// Get User session to check if admin & set IsAdmin
	var user Session_User
	var isAdmin bool
	u := this.GetSession("_User")
	if u != nil {
		user = u.(Session_User)
		this.Data["User"] = user
		if user.Username == "BTCSport" {
			isAdmin = true
		}
	}
	this.Data["IsAdmin"] = isAdmin

	reqUrl := this.Ctx.Request.URL.String()
	sec := reqUrl[strings.LastIndex(reqUrl, "/")+1:]
	if qm := strings.Index(sec, "?"); qm > -1 {
		sec = sec[:qm]
	}

	if len(sec) == 0 || sec == "admin" {
		sec = "User_All"
		this.Data[sec] = true
	} else {
		this.Data[sec] = true
	}

	this.setData(sec)
}

func (this *AdminRouter) setData(sec string) {
	switch sec {
	case "User_All":
		this.Data["UserAll"] = models.AllUsers()
	case "User_Refer":
		this.Data["UserAll"] = models.AllUsers()
	case "Game_New":
		this.Data["xsrf"] = template.HTML(this.XsrfFormHtml())
	case "Game_Pending":
		this.Data["xsrf"] = template.HTML(this.XsrfFormHtml())
		this.Data["GamePending"] = models.AllGamePending()
	case "Game_End":
		this.Data["GameEnded"] = models.AllGameEnded()
	case "Wallet_Deposit":
		this.Data["Deposit"] = models.AllDeposit()
	case "Wallet_Withdraw":
		this.Data["Withdraw"] = models.AllWithdraw()
	case "Wallet_Pay":

	}
	this.TplNames = "admin_" + sec + ".tpl"
}

func (this *AdminRouter) Post() {

	// Get User session to check if admin & set IsAdmin
	var user Session_User
	var isAdmin bool
	u := this.GetSession("_User")
	if u != nil {
		user = u.(Session_User)
		if user.Username == "BTCSport" {
			isAdmin = true
		}
	}
	isAdmin = true

	if !isAdmin {
		this.Ctx.Redirect(302, "/admin")
		return
	}

	reqUrl := this.Ctx.Request.URL.String()
	sec := reqUrl[strings.LastIndex(reqUrl, "/")+1:]
	if qm := strings.Index(sec, "?"); qm > -1 {
		sec = sec[:qm]
	}

	this.postData(sec)
}

func (this *AdminRouter) postData(sec string) {
	inputs := this.Input()
	switch sec {
	case "Game_New":
		teamA := inputs.Get("TeamA")
		teamB := inputs.Get("TeamB")
		oddsa := inputs.Get("Oddsa")
		oddsb := inputs.Get("Oddsb")
		concede := inputs.Get("Concede")
		scoresum := inputs.Get("ScoreSum")
		starttime := inputs.Get("TimeStart")
		Type := inputs.Get("Type")
		slice := []interface{}{"Football", "Basketball"}

		// Check if input matches
		if models.ValidOdds(oddsa) && models.ValidOdds(oddsb) && models.ValidScore(concede) && models.ValidScore(scoresum) && models.ValidStarttime(starttime) && In_slice(Type, slice) {
			oddsa_float64, _ := strconv.ParseFloat(oddsa, 64)
			oddsb_float64, _ := strconv.ParseFloat(oddsb, 64)
			concede_float64, _ := strconv.ParseFloat(concede, 64)
			scoresum_float64, _ := strconv.ParseFloat(scoresum, 64)
			var isfootball int
			if Type == "Football" {
				isfootball = 1
			}

			if !models.NewGame(models.Game_Detail{Isfootball: isfootball, Teama: teamA, Teamb: teamB, Oddsa: oddsa_float64, Oddsb: oddsb_float64, Concede: concede_float64, Scoresum: scoresum_float64, Timestarted: starttime, Timecreated: time.Now().Format(layout)}) {
				models.Log(models.Log_Struct{"error", "Create New Game:", errors.New("Failed to create new game")})
			}
		}
	case "Game_Result":
		gid := inputs.Get("gid")
		result := inputs.Get("result")

		// Check if input matches
		if models.ValidResult(result) && models.ValidGid(gid) {
			r := strings.Split(result, ":")
			score_a, _ := strconv.Atoi(r[0])
			score_b, _ := strconv.Atoi(r[1])
			score_sum := score_a + score_b

			id, _ := strconv.Atoi(gid)

			// Update the game
			if !models.UpdateGameById(score_a, score_b, id) {
				panic("Cant Update Result")
			}

			// Get the game by id, get game scoresum, odds, concede
			game := models.GameById(id)
			sum := game.Scoresum
			oddsa := game.Oddsa
			oddsb := game.Oddsb
			concede := game.Concede

			// Calculate the result, A_Win, B_Win, Odd, Even, Large, Small
			odds := oddsa
			a_or_b := "A_Win"
			if score_a-int(concede) <= score_b {
				a_or_b = "B_Win"
				odds = oddsb
			}
			oddeven := "Odd"
			if score_sum%2 == 0 {
				oddeven = "Even"
			}
			largesmall := "Large"
			if score_sum <= int(sum) {
				largesmall = "Small"
			}

			g_result := []interface{}{a_or_b, oddeven, largesmall}
			// Update_Distribution, TODO:
			if !models.CalculateResult(id, odds, ProfitAddr, g_result) {
				panic("Cant Calculate Result")
			}
		}
	}
	this.Ctx.Redirect(302, "/admin")
}

func In_slice(val interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
