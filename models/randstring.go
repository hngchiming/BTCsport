package models

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	defaultRand = rand.New(rand.NewSource(time.Now().Add(10086 * time.Minute).UnixNano()))
)

func RandString(l int) string {
	h := hmac.New(sha512.New, []byte(fmt.Sprintf("%d", defaultRand.Int())))
	h.Write([]byte("Wo@iLWS1314"))
	sign := hex.EncodeToString(h.Sum(nil))
	r := []rune(sign)
	var temp []rune
	temp = append(temp, r[:l]...)
	return strings.ToUpper(string(temp))
}
