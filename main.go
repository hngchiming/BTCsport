package main

import (
	"BTCsport/controllers"
	"github.com/astaxie/beego"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	beego.SessionOn = true
	beego.SessionName = "BTCSport"
	beego.HttpPort = 8080

	os.Mkdir("./log", os.ModePerm)
	fw := beego.NewFileWriter("./log/log", true)
	fw.SetRotateMaxDays(30)
	fw.SetRotateSize(1 << 25)
	err := fw.StartLogger()
	if err != nil {
		beego.Critical("NewFileWriter ->", err)
	}

	beego.EnableXSRF = true
	beego.XSRFKEY = "WO@ilws1314"
	beego.XSRFExpire = 300
}

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginRouter{})
	beego.Router("/register", &controllers.RegisterRouter{})
	beego.Router("/forget", &controllers.ForgetRouter{})
	beego.Router("/reset", &controllers.ResetRouter{})
	beego.Router("/football", &controllers.FootballRouter{})
	beego.Router("/basketball", &controllers.BasketballRouter{})
	beego.Router("/game/:id:int", &controllers.GameRouter{})

	// About
	beego.Router("/about", &controllers.AboutRouter{})

	// Chat
	beego.Router("/chat", &controllers.ChatRouter{})
	beego.RouterHandler("/chat/:info(.*)", controllers.SockjsHandler)

	// Withdraw
	beego.Router("/withdraw", &controllers.WithdrawRouter{})

	// Userinfo
	beego.Router("/user", &controllers.UserRouter{})
	beego.Router("/user/:all", &controllers.UserRouter{})

	// Admin
	beego.Router("/admin", &controllers.AdminRouter{})
	beego.Router("/admin/:all", &controllers.AdminRouter{})

	// 404
	beego.Router("/:all", &controllers.ErrorRouter{})

	beego.Run()
}
