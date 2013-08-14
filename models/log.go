package models

import (
	"github.com/astaxie/beego"
)

type funcMap map[string]func(...interface{})
type funcSlice []interface{}

type Log_Struct struct {
	Logtype, Logstring string
	Err                error
}

var (
	m funcMap
	s funcSlice
)

func init() {
	m = make(funcMap)
	m = funcMap{
		"trace":    beego.Trace,
		"debug":    beego.Debug,
		"info":     beego.Info,
		"warn":     beego.Warn,
		"error":    beego.Error,
		"critical": beego.Critical,
	}
	s = make(funcSlice, 0, 6)
	s = append(s, funcSlice{
		"trace",
		"debug",
		"info",
		"warn",
		"error",
		"critical"}...)
}

func in_slice(val interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func Log(ls Log_Struct) {
	if !in_slice(ls.Logtype, s) {
		ls.Logtype = "trace"
	}
	m[ls.Logtype](ls.Logstring, ls.Err)
}
