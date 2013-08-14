package models

import (
	"errors"
	"fmt"
	"time"
)

const (
	layout    = "2006-01-02 15:04:05"
	app_start = "2013-08-30 00:00:00"
	tx_fee    = 0.0005
	db_front  = "./test.db"
	email     = "bitcoinsport"
	emailpass = "03RQ7dVE1J"
	appname   = "BTC体育竞猜"
)

var checkTicker *time.Ticker

func init() {
	// Check deposit every 5 minutes
	checkTicker = time.NewTicker(5 * time.Minute)
	go checkTickerTimer(checkTicker.C)

	//checkDeposit()
}

func checkTickerTimer(checkchan <-chan time.Time) {
	for {
		<-checkchan
		checkDeposit()
	}
}

func checkDeposit() {
	allUsers := AllUsers()
	for _, v := range allUsers {
		bal, err := GetBalance(v.Btcaddress)
		if err != nil {
			Log(Log_Struct{"critical", "Deposit_Checking:", err})
			fmt.Println(err)
			return
		}

		if bal > v.Balance {
			amount := bal - v.Balance
			if !UserDeposit(v.Id, amount) {
				Log(Log_Struct{"critical", "Deposit_Checking:", errors.New("DB_Maybe_Locked")})
				fmt.Println("DB_Maybe_Locked")
				return
			}
		}
	}
}
