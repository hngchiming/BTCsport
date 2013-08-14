package controllers

import (
	"BTCsport/models"
	"github.com/astaxie/beego"
	"html/template"
)

type FootballRouter struct {
	beego.Controller
}

func (this *FootballRouter) Get() {
	// set const
	this.Data["App_Name"] = App_Name

	this.Data["IsFootball"] = true

	// Get User Session
	var user Session_User
	u := this.GetSession("_User")
	if u != nil {
		user = u.(Session_User)
		this.Data["User"] = user
	} else {
		this.Data["User"] = false
	}

	// Get all Footballball games
	all := models.AllFootball()
	var AllFootball []Game
	for _, v := range all {

		// Check if game is over
		if v.Isover == 1 {
			AllFootball = append(AllFootball, Game{
				Game_Detail: v,
				Tr_Class:    template.HTML("error"),
				Lbl_Class:   template.HTML("-important"),
				Btn_Class:   template.HTML("-danger"),
				Lbl_String:  "已结束",
				Btn_String:  "查看详情",
			})
		}

		// Starttime sub now
		timeSub := TimeSub(v.Timestarted).Minutes()

		switch {
		case timeSub < 0:
			AllFootball = append(AllFootball, Game{
				Game_Detail: v,
				Tr_Class:    template.HTML("warning"),
				Lbl_Class:   template.HTML("-warning"),
				Btn_Class:   template.HTML("-danger"),
				Lbl_String:  "正在进行中",
				Btn_String:  "查看详情",
			})
		case timeSub < 15:
			AllFootball = append(AllFootball, Game{
				Game_Detail: v,
				Tr_Class:    template.HTML("warning"),
				Lbl_Class:   template.HTML(""),
				Btn_Class:   template.HTML("-danger"),
				Lbl_String:  "比赛即将开始",
				Btn_String:  "查看详情",
			})
		default:
			AllFootball = append(AllFootball, Game{
				Game_Detail: v,
				Tr_Class:    template.HTML("success"),
				Lbl_Class:   template.HTML("-success"),
				Btn_Class:   template.HTML("-success"),
				Lbl_String:  "可参与",
				Btn_String:  "参与",
			})
		}
	}

	this.Data["AllGame"] = AllFootball

	this.TplNames = "football.html"
}
