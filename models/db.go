package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/mattn/go-sqlite3"
	"reflect"
	"time"
)

type User struct {
	Id           int
	Username     string
	Password     string
	Fundpassword string
	Email        string
	Birth        string
	Btcaddress   string
	Balance      float64
	Profit       float64
	Alltimebet   float64
	Lastip       string
	Referral     string
}

type User_Bet struct {
	Uid       int
	Gid       int
	Type      string
	Betamount float64
	Txhash    string
	Bettime   string
	Profit    float64
	Txhashwin string
}

type Game_Detail struct {
	Id          int
	Isfootball  int
	Teama       string
	Teamb       string
	Scorea      int
	Scoreb      int
	Oddsa       float64
	Oddsb       float64
	Concede     float64
	Scoresum    float64
	Timestarted string
	Timecreated string
	Poolwin     float64
	Poollose    float64
	Pooleven    float64
	Poolodd     float64
	Poollarge   float64
	Poolsmall   float64
	Poolsum     float64
	Isover      int
}

type Current_Bet struct {
	Gid    int
	Uid    int
	Txhash string
	Type   string
	Bet    float64
}

type Deposit struct {
	Uid    int
	Amount float64
	Time   string
}

type Withdraw struct {
	Uid     int
	Amount  float64
	Address string
	Time    string
}

type Gamehistory struct {
	Id         int
	Siteprofit float64
}

func check_err(err error) bool {
	if err != nil {
		Log(Log_Struct{"error", "DB_Error:", err})
		return false
	}
	return true
}

/*
	Get orm Model
*/
func get_DBFront() beedb.Model {
	db, err := sql.Open("sqlite3", db_front)
	if !check_err(err) {
		panic(err)
	}
	orm := beedb.New(db)
	return orm
}

/*
	Add New User To DB: return bool
*/
func NewUser(u User) bool {
	orm := get_DBFront()
	err := orm.SetTable("user").Save(&u)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_115", err})
		return false
	}
	return true
}

/*
	Get User by uid
*/
func UserByUid(uid int) User {
	orm := get_DBFront()
	var user User
	err := orm.SetTable("user").Where(uid).Find(&user)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_129", err})
	}
	return user
}

/*
	Get Uid by Username
*/
func UidByUsername(username string) int {
	orm := get_DBFront()
	var user User
	err := orm.SetTable("user").Where("username=?", username).Find(&user)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_142", err})
		return 0
	}

	return user.Id
}

/*
	Get User By Referral
*/
func UserByReferral(referral string) []User {
	orm := get_DBFront()
	var user []User
	err := orm.SetTable("user").Where("referral=?", referral).FindAll(&user)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_157", err})
	}
	return user
}

/*
	Check if enough balance, return bool
*/
func UserBalanceEnough(uid int, amount float64) bool {
	orm := get_DBFront()
	var user User
	err := orm.SetTable("user").Where(uid).Find(&user)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_170", err})
		return false
	}

	return user.Balance >= amount
}

/*
	Check if User exists, return IsExist
*/
func UserExist(username string) bool {
	orm := get_DBFront()
	var user User
	err := orm.SetTable("user").Where("username=?", username).Find(&user)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_185", err})
		return false
	}
	return true
}

/*
	Check if Email exists, return IsExist
*/
func EmailExist(email string) bool {
	orm := get_DBFront()
	var user User
	err := orm.SetTable("user").Where("email=?", email).Find(&user)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_199", err})
		return false
	}
	return true
}

/*
	Check Birth of User matches the string or not, return bool
*/
func BirthMatch(username, birth string) bool {
	orm := get_DBFront()
	var user User
	err := orm.SetTable("user").Where("username=?", username).Find(&user)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_213", err})
		return false
	}

	if user.Birth == birth {
		return true
	}
	return false
}

/*
	Check Email of User matches the string or not, return bool
*/
func EmailMatch(username, email string) bool {
	orm := get_DBFront()
	var user User
	err := orm.SetTable("user").Where("username=?", username).Find(&user)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_231", err})
		return false
	}

	if user.Email == email {
		return true
	}
	return false
}

/*
	Update Password and Fundpassword of a user
*/
func UpdateUserPass(username, pass, fundpass string) bool {
	orm := get_DBFront()
	t := make(map[string]interface{})
	t["password"] = pass
	t["fundpassword"] = fundpass
	_, err := orm.SetTable("user").Where("username=?", username).Update(t)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_251", err})
		return false
	}
	return true
}

/*
	Check if password matches, return bool
*/
func PassMatch(username, password string) bool {
	orm := get_DBFront()
	var user User
	err := orm.SetTable("user").Where("username=?", username).Find(&user)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_265", err})
		return false
	}

	return match_pass(password, user.Password)
}

/*
	Check if fundpassword matches, return bool
*/
func FundPassMatch(username, password string) bool {
	orm := get_DBFront()
	var user User
	err := orm.SetTable("user").Where("username=?", username).Find(&user)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_280", err})
		return false
	}

	return match_pass(password, user.Fundpassword)
}

/*
	Get User BTC address
*/
func AddressByUid(uid int) string {
	orm := get_DBFront()
	var user User
	err := orm.SetTable("user").Where(uid).Find(&user)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_295", err})
	}
	return user.Btcaddress
}

/*
	Update user last ip
*/
func UpdateIP(username, ip string) {
	orm := get_DBFront()
	t := make(map[string]interface{})
	t["lastip"] = ip
	_, err := orm.SetTable("user").Where("username=?", username).Update(t)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_309", err})
	}
}

/*
	Get all users for admin
*/
func AllUsers() []User {
	orm := get_DBFront()
	var allUser []User
	err := orm.SetTable("user").FindAll(&allUser)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_321", err})
	}
	return allUser
}

/*
	Get all pending games
*/
func AllGamePending() []Game_Detail {
	orm := get_DBFront()
	var allGame, allPendingGame []Game_Detail
	err := orm.SetTable("game").FindAll(&allGame)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_334", err})
		return allGame
	}

	for _, v := range allGame {
		startTime, _ := time.Parse(layout, v.Timestarted)
		timeNow, _ := time.Parse(layout, time.Now().String())
		if startTime.Sub(timeNow) > 15*time.Minute {
			allPendingGame = append(allPendingGame, v)
		}
	}

	SliceReverse(allPendingGame)
	return allPendingGame
}

/*
	Get all end games
*/
func AllGameEnded() []Game_Detail {
	orm := get_DBFront()
	var allGame, allEndedGame []Game_Detail
	err := orm.SetTable("game").FindAll(&allGame)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_358", err})
		return allGame
	}

	for _, v := range allGame {
		if v.Isover == 1 {
			allEndedGame = append(allEndedGame, v)
		}
	}

	SliceReverse(allEndedGame)
	return allEndedGame
}

/*
	Get all basketball
*/
func AllBasketball() []Game_Detail {
	orm := get_DBFront()
	var allGame, allBasketball []Game_Detail
	err := orm.SetTable("game").FindAll(&allGame)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_380", err})
		return allGame
	}

	for _, v := range allGame {
		if v.Isfootball != 1 {
			allBasketball = append(allBasketball, v)
		}
	}

	SliceReverse(allBasketball)
	return allBasketball
}

/*
	Get all football
*/
func AllFootball() []Game_Detail {
	orm := get_DBFront()
	var allGame, allFootball []Game_Detail
	err := orm.SetTable("game").FindAll(&allGame)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_402", err})
		return allGame
	}

	for _, v := range allGame {
		if v.Isfootball == 1 {
			allFootball = append(allFootball, v)
		}
	}

	SliceReverse(allFootball)
	return allFootball
}

/*
	New game, return bool
*/
func NewGame(game Game_Detail) bool {
	orm := get_DBFront()
	err := orm.SetTable("game").Save(&game)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_423", err})
		return false
	}
	return true
}

/*
	Get game by id
*/
func GameById(id int) Game_Detail {
	orm := get_DBFront()
	var game Game_Detail
	err := orm.SetTable("game").Where(id).Find(&game)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_437", err})
	}

	return game
}

/*
	Update game by id
*/
func UpdateGameById(a, b, id int) bool {
	orm := get_DBFront()
	t := make(map[string]interface{})
	t["scorea"] = a
	t["scoreb"] = b
	_, err := orm.SetTable("game").Where(id).Update(t)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_453", err})
		return false
	}
	return true
}

/*
	Place a bet, return bool
*/
func PlaceBet(gid, uid int, t, txhash string, amount float64) bool {
	db, err := sql.Open("sqlite3", db_front)
	if !check_err(err) {
		return false
	}
	defer db.Close()

	var query string
	tx, err := db.Begin()
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_472", err})
		return false
	}

	// Insert
	query = "INSERT INTO currentbet(txhash, gid, uid, type, bet) values(?,?,?,?,?)"
	stmt, err := tx.Prepare(query)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_480", err})
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(txhash, gid, uid, t, amount-tx_fee)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_486", err})
		return false
	}

	// Update Game, pool poolsum, amount = amount - transaction fee 0.0001
	switch t {
	case "A_Win":
		query = "UPDATE game SET poolwin = poolwin + %f, poolsum = poolsum + %f WHERE id = %d"
	case "B_Win":
		query = "UPDATE game SET poollose = poollose + %f, poolsum = poolsum + %f WHERE id = %d"
	case "Odd":
		query = "UPDATE game SET poolodd = poolodd + %f, poolsum = poolsum + %f WHERE id = %d"
	case "Even":
		query = "UPDATE game SET pooleven = pooleven + %f, poolsum = poolsum + %f WHERE id = %d"
	case "Large":
		query = "UPDATE game SET poollarge = poollarge + %f, poolsum = poolsum + %f WHERE id = %d"
	case "Small":
		query = "UPDATE game SET poolsmall = poolsmall + %f, poolsum = poolsum + %f WHERE id = %d"
	}
	query = fmt.Sprintf(query, amount-tx_fee, amount-tx_fee, gid)
	stmt, err = tx.Prepare(query)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_508", err})
		return false
	}
	_, err = stmt.Exec()
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_513", err})
		return false
	}

	// Update User, set balance = balance - amount
	query = fmt.Sprintf("UPDATE user SET balance = balance - %f, alltimebet = alltimebet + %f WHERE id = %d", amount, amount-tx_fee, uid)
	stmt, err = tx.Prepare(query)
	_, err = stmt.Exec()
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_522", err})
		return false
	}

	// Insert alluserbet, set betamount = betamount + amount - tx_fee
	query = fmt.Sprintf("INSERT INTO alluserbet(uid, gid, type, betamount,txhash, bettime) values(?,?,?,?,?,?)")
	stmt, err = tx.Prepare(query)
	_, err = stmt.Exec(uid, gid, t, amount-tx_fee, txhash, time.Now().Format(layout))
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_531", err})
		return false
	}

	err = tx.Commit()
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_537", err})
		tx.Rollback()
		return false
	}

	return true
}

/*
	User Withdraw BTC, need to update balance and record it in withdraw
*/
func WithdrawRequest(uid int, amount float64, address string) bool {
	db, err := sql.Open("sqlite3", db_front)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_551", err})
		return false
	}
	defer db.Close()

	var query string
	tx, err := db.Begin()
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_559", err})
		return false
	}

	// Insert into withdraw record
	query = "INSERT INTO withdraw(uid,amount,address,time) values(?,?,?,?)"
	stmt, err := tx.Prepare(query)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_567", err})
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, amount, address, time.Now().Format(layout))
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_573", err})
		return false
	}

	// Update user info
	query = fmt.Sprintf("UPDATE user SET balance = balance - %f WHERE id = %d", amount, uid)
	stmt, err = tx.Prepare(query)
	_, err = stmt.Exec()
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_582", err})
		return false
	}

	// Commit
	err = tx.Commit()
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_589", err})
		tx.Rollback()
		return false
	}

	return true

}

/*
	User deposit, need to update User and insert into deposit record
*/
func UserDeposit(uid int, amount float64) bool {
	db, err := sql.Open("sqlite3", db_front)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_604", err})
		return false
	}
	defer db.Close()

	var query string
	tx, err := db.Begin()
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_612", err})
		return false
	}

	// Insert into deposit record
	query = "INSERT INTO deposit(uid,amount,time) values(?,?,?)"
	stmt, err := tx.Prepare(query)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_620", err})
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, amount, time.Now().Format(layout))
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_626", err})
		return false
	}

	// Update user info
	query = fmt.Sprintf("UPDATE user SET balance = balance + %f WHERE id = %d", amount, uid)
	stmt, err = tx.Prepare(query)
	_, err = stmt.Exec()
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_635", err})
		return false
	}

	// Commit
	err = tx.Commit()
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_642", err})
		tx.Rollback()
		return false
	}

	return true
}

/*
	All Current Bet
*/
func AllCurBet() []Current_Bet {
	orm := get_DBFront()
	var cur []Current_Bet
	err := orm.SetTable("currentbet").FindAll(&cur)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_658", err})
	}
	return cur
}

/*
	Get all Deposit, return slice
*/
func AllDeposit() []Deposit {
	orm := get_DBFront()
	var depo []Deposit
	err := orm.SetTable("deposit").FindAll(&depo)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_671", err})
	}
	SliceReverse(depo)
	return depo
}

/*
	Get Deposit by Uid, return slice
*/
func DepositByUid(uid int) []Deposit {
	var depo []Deposit
	for _, v := range AllDeposit() {
		if v.Uid == uid {
			depo = append(depo, v)
		}
	}
	SliceReverse(depo)
	return depo
}

/*
	Get all Withdraw, return slice
*/
func AllWithdraw() []Withdraw {
	orm := get_DBFront()
	var with []Withdraw
	err := orm.SetTable("withdraw").FindAll(&with)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_699", err})
	}
	SliceReverse(with)
	return with
}

/*
	Get Withdraw by Uid, return slice
*/
func WithdrawByUid(uid int) []Withdraw {
	var with []Withdraw
	for _, v := range AllWithdraw() {
		if v.Uid == uid {
			with = append(with, v)
		}
	}
	SliceReverse(with)
	return with
}

/*
	Get All User Bet: return slice
*/
func AllUserBets() []User_Bet {
	orm := get_DBFront()
	var bets []User_Bet
	err := orm.SetTable("alluserbet").FindAll(&bets)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_727", err})
	}
	return bets
}

/*
	Get User Bet by Uid
*/
func UserBets(uid int) []User_Bet {
	orm := get_DBFront()
	var bets []User_Bet
	err := orm.SetTable("alluserbet").Where("uid=?", uid).FindAll(&bets)
	if !check_err(err) {
		Log(Log_Struct{"error", "DB_Error_Line_740", err})
	}
	SliceReverse(bets)
	return bets
}

/*
	Get user bet amount and profit of this month
*/
func UserBetAmountProfit(uid int) (float64, float64) {
	var amount, profit float64
	bets := UserBets(uid)
	for _, v := range bets {
		bettime, _ := time.Parse(layout, v.Bettime)
		appstarttime, _ := time.Parse(layout, app_start)
		temp := appstarttime.AddDate(0, Month()-1, 0)
		if bettime.After(temp) {
			amount += v.Betamount
			profit += v.Profit
		}
	}
	return amount, profit
}

/*
	Reverse a slice
*/
func SliceReverse(a interface{}) {
	rv := reflect.ValueOf(a)

	if !rv.IsValid() {
		return
	}

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return
	}

	for i := 0; i < rv.Len()/2; i++ {
		temp := rv.Index(rv.Len() - i - 1).Interface()
		rv.Index(rv.Len() - i - 1).Set(rv.Index(i))
		rv.Index(i).Set(reflect.ValueOf(temp))
	}
}

/*
	Get Month since Appstart
*/
func Month() int {
	t_start, _ := time.Parse(layout, app_start)
	t_now, _ := time.Parse(layout, time.Now().Format(layout))
	i := 0
	for {
		i += 1
		if t_start.AddDate(0, i, 0).After(t_now) {
			return i
		}
	}
	return 0
}
