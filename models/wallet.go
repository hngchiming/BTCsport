package models

import (
	"errors"
	"github.com/astaxie/beego/httplib"
	"strconv"
	"time"
)

const (
	satoshi    = 100000000
	wallet_url = "https://blockchain.info/merchant/a3ff8c3c-e8ed-4daf-ad20-a8fb29b510d4/"
	firstpass  = "Hu@ngwen41314"
	secondpass = "BnProMiAch"
	passwords  = "password=Hu@ngwen41314&second_password=BnProMiAch"
)

type error_callback struct {
	Error string `json:"error"`
}

type address_callback struct {
	Address string `json:"address"`
	Label   string `json:"label"`
}

/*
	Generate new address for new user
	address, err := models.NewAddress(username)
	fmt.Println(address)
*/
func NewAddress(username string) (string, error) {
	url := wallet_url
	method := "new_address?"
	pas := passwords
	body := httplib.Get(url+method+pas+"&label="+username).SetTimeout(3*time.Second, 2*time.Second)

	var callback address_callback
	err := body.ToJson(&callback)
	if err != nil {
		return "", err
	}

	if callback.Address == "" {
		var err_callback error_callback
		err = body.ToJson(&err_callback)
		if err != nil {
			return "", err
		}
		return "", errors.New(err_callback.Error)
	}

	return callback.Address, nil
}

type getbalance_callback struct {
	Balance float64 `json:"balance"`
	Address string  `json:"address"`
	Total   float64 `json:"total_received"`
}

/*
	Get balance of specific address
	bal, err := models.GetBalance("1HnhxpDg4FzK5wYkyJH4NNVoq4B6FkbsLZ")
	fmt.Println(bal)
*/
func GetBalance(address string) (float64, error) {
	url := wallet_url
	method := "address_balance?"
	pas := firstpass
	body := httplib.Get(url+method+"password="+pas+"&address="+address+"&confirmations=2").SetTimeout(3*time.Second, 2*time.Second)
	var callback getbalance_callback
	err := body.ToJson(&callback)
	if err != nil {
		return 0.0, err
	}

	if callback.Address == "" {
		var err_callback error_callback
		err = body.ToJson(&err_callback)
		if err != nil {
			return 0.0, err
		}
		return 0.0, errors.New(err_callback.Error)
	}

	return callback.Balance / satoshi, nil
}

type List struct {
	AllAddress []Alladdress_callback `json:"addresses"`
}

type Alladdress_callback struct {
	Balance float64 `json:"balance"`
	Address string  `json:"address"`
	Label   string  `json:"label"`
	Total   float64 `json:"total_received"`
}

/*
	List all the address and balance --> for update
	addSlice, err := models.ListAddress()
	fmt.Println(addSlice)
*/
func ListAddress() ([]Alladdress_callback, error) {
	url := wallet_url
	method := "list?"
	pas := firstpass
	body := httplib.Get(url+method+"password="+pas).SetTimeout(3*time.Second, 2*time.Second)
	var callback List
	err := body.ToJson(&callback)
	if err != nil {
		return []Alladdress_callback{}, err
	}

	if len(callback.AllAddress) == 0 {
		var err_callback error_callback
		err := body.ToJson(&err_callback)
		if err != nil {
			return []Alladdress_callback{}, err
		}
		return []Alladdress_callback{}, errors.New(err_callback.Error)
	}

	for i, v := range callback.AllAddress {
		callback.AllAddress[i].Balance = v.Balance / satoshi
	}
	return callback.AllAddress, nil
}

type send_callback struct {
	Message string `json:"message"`
	TX_Hash string `json:"tx_hash"`
	Notice  string `json:"notice"`
}

/*
 Send an Amount of btc From address To address with Note
 tx_hash, err := models.SendBTC("1A3wRZGBPKg6cTPYbo5MqarvU9cfpEAnfc", "1ALP9rrweYEksfbspup1Fi6w6QcGp7Z7Fq", "Win", 0.00010001)
 fmt.Println(tx_hash)
*/
func SendBTC(from, to, note string, amount float64) (string, error) {
	url := wallet_url
	method := "payment?"
	pas := passwords
	body := httplib.Get(url+method+pas+"&from="+from+"&to="+to+"&amount="+strconv.FormatFloat(amount*satoshi, 'f', -1, 64)+"&note="+note).SetTimeout(3*time.Second, 2*time.Second)
	var callback send_callback
	err := body.ToJson(&callback)
	if err != nil {
		return "", err
	}

	if callback.Message == "" {
		var err_callback error_callback
		err := body.ToJson(&err_callback)
		if err != nil {
			return "", err
		}
		return "", errors.New(err_callback.Error)
	}

	return callback.TX_Hash, nil
}

func Archive(address string) error {
	url := wallet_url
	method := "archive_address?"
	pas := passwords
	body := httplib.Get(url+method+pas+"&address="+address).SetTimeout(3*time.Second, 2*time.Second)
	var callback error_callback
	err := body.ToJson(&callback)
	if err != nil {
		return err
	}

	if callback.Error != "" {
		return errors.New(callback.Error)
	}

	return nil
}
