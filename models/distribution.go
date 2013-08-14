package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

/*
	Calculate the result and update/insert/delete something, THE MOST IMPORTANT FUNCTION
*/
func CalculateResult(gid int, odds float64, profitAddr string, result []interface{}) bool {

	// Get UserBet by gid
	var userBet []User_Bet
	for _, v := range AllUserBets() {
		if v.Gid == gid {
			userBet = append(userBet, v)
		}
	}

	// Loop ->
	// 1. Check result, get payback, profit ->
	// 2. Send BTC profit ->
	// transaction start
	// 3. Update user balance + payback, profit + profit ->
	// 5. Update alluserbet profit + profit, txhashwin ->
	// 6. Insert into historybet alluserbet ->
	// 7. Delete from current bet ->
	// transaction end -> Loop

	// Slice []interface{}
	abSlice := []interface{}{"A_Win", "B_Win"}
	for _, v := range userBet {
		// define error
		var err error

		// 1. Check result, get payback, profit
		var payback, profit float64
		if in_slice(v.Type, result) {
			if in_slice(v.Type, abSlice) {
				payback = v.Betamount*odds - 0.0005
				profit = v.Betamount*(odds-1) - 0.0005
			} else {
				payback = v.Betamount*1.95 - 0.0005
				profit = v.Betamount*0.95 - 0.0005
			}
		} else {
			payback = 0
			profit = -v.Betamount
		}

		// 2. Get user address and Send BTC if profit > 0
		addr := UserByUid(v.Uid).Btcaddress

		var txhashwin string
		if profit > 0 {
			txhashwin, err = SendBTC(profitAddr, addr, "竞猜胜利获得奖励："+fmt.Sprintf("%f", payback), payback)
			if err != nil {
				Log(Log_Struct{"error", "Calculate Result:", err})
				return false
			}
		}

		// StartTransaction
		db, err := sql.Open("sqlite3", db_front)
		if !check_err(err) {
			return false
		}
		defer db.Close()

		var query string
		tx, err := db.Begin()
		if !check_err(err) {
			return false
		}

		// 3. Update user balance, profit
		query = fmt.Sprintf("UPDATE user SET balance = balance + %f, profit = profit + %f WHERE id = %d", payback, profit, v.Uid)
		stmt, err := tx.Prepare(query)
		if !check_err(err) {
			return false
		}
		defer stmt.Close()
		_, err = stmt.Exec()
		if !check_err(err) {
			return false
		}

		// 5. Update alluserbet, profit and txhashwin
		query = fmt.Sprintf("UPDATE alluserbet SET profit = profit + %f WHERE txhash = %s", profit, v.Txhash)
		stmt, err = tx.Prepare(query)
		if !check_err(err) {
			return false
		}
		defer stmt.Close()
		_, err = stmt.Exec()
		if !check_err(err) {
			return false
		}

		// 6. Insert into historybet
		query = "INSERT INTO historybet(uid, gid, type, betamount, txhash, bettime, profit, txhashwin) values(?,?,?,?,?,?,?,?)"
		stmt, err = tx.Prepare(query)
		if !check_err(err) {
			return false
		}
		defer stmt.Close()
		_, err = stmt.Exec(v.Uid, v.Gid, v.Type, v.Betamount, v.Txhash, v.Bettime, profit, txhashwin)
		if !check_err(err) {
			return false
		}

		// 7. Delete from currentbet
		query = fmt.Sprintf("DELETE FROM currentbet WHERE txhash = %s", v.Txhash)
		stmt, err = tx.Prepare(query)
		if !check_err(err) {
			return false
		}
		defer stmt.Close()
		_, err = stmt.Exec()
		if !check_err(err) {
			return false
		}

		// transaction close
		err = tx.Commit()
		if !check_err(err) {
			tx.Rollback()
			return false
		}

		db.Close()
	}

	return true
}
