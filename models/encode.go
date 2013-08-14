package models

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func EncodePass(pass string) string {
	h := hmac.New(sha512.New, []byte(pass))
	h.Write([]byte("Wo@iLWS1314"))
	return hex.EncodeToString(h.Sum(nil))
}

func match_pass(pass, encodepass string) bool {
	return encodepass == EncodePass(pass)
}
