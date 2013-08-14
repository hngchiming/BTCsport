package controllers

import (
	"BTCsport/models"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/astaxie/beego"
	"html/template"
	"time"
)

const (
	Tx_Fee   = 0.0005
	App_Name = "BTC体育竞猜"
	layout   = "2006-01-02 15:04:05"
)

const (
	ERROR_USEREXIST = iota
	ERROR_USERNOTEXIST
	ERROR_BIRTHNOTMATCH
	ERROR_EMAILNOTMATCH
	ERROR_EMAILNOTSENT
	ERROR_EMAILEXIST
	ERROR_ADDRESS
	ERROR_DB
	ERROR_INVALIDINPUT
	ERROR_TWICESUBMIT
	ERROR_CANTUPDATE
	ERROR_CODENOTMATCH
	ERROR_CAPTCHA
	ERROR_PASSINCORRECT
	ERROR_BALANCENOTENOUGH
	ERROR_CANTSENDBTC
	ERROR_REFERNOTEXIST
)

var (
	BetAddr    string
	ProfitAddr string
)

func init() {
	BetAddr = beego.AppConfig.String("BetAddr")
	ProfitAddr = beego.AppConfig.String("ProfitAddr")
}

type Session_User struct {
	Uid      int
	Username string
	Ip       string
}

type Game struct {
	Game_Detail models.Game_Detail
	Tr_Class    template.HTML
	Lbl_Class   template.HTML
	Btn_Class   template.HTML
	Lbl_String  template.HTML
	Btn_String  template.HTML
}

func TimeSub(time_str string) time.Duration {
	t_start, _ := time.Parse(layout, time_str)
	t_now, _ := time.Parse(layout, time.Now().Format(layout))
	return t_start.Sub(t_now)
}

func GetError(e interface{}) string {
	switch e.(int) {
	case ERROR_USEREXIST:
		return "该用户名已经被注册"
	case ERROR_USERNOTEXIST:
		return "该用户不存在，请刷新后重试"
	case ERROR_BIRTHNOTMATCH:
		return "出生日期不正确，请重试"
	case ERROR_EMAILNOTMATCH:
		return "邮箱地址不正确"
	case ERROR_EMAILNOTSENT:
		return "无法发送验证码，请稍后重试"
	case ERROR_EMAILEXIST:
		return "该邮箱已经被注册"
	case ERROR_ADDRESS:
		return "无法生成新地址，请稍后重试"
	case ERROR_DB:
		return "系统繁忙...请重试"
	case ERROR_INVALIDINPUT:
		return "输入的信息有误，请按照格式重试"
	case ERROR_TWICESUBMIT:
		return "请勿重复提交请求，刷新后重试"
	case ERROR_CANTUPDATE:
		return "系统繁忙，无法重设密码，请稍后再试"
	case ERROR_CODENOTMATCH:
		return "验证码错误，请前往注册的邮箱获取验证码"
	case ERROR_CAPTCHA:
		return "验证码错误，拒绝机器人~"
	case ERROR_PASSINCORRECT:
		return "密码错误，请刷新后重试"
	case ERROR_BALANCENOTENOUGH:
		return "余额不足，请先充值~"
	case ERROR_CANTSENDBTC:
		return "无法发送比特币，请稍后重试"
	case ERROR_REFERNOTEXIST:
		return "推荐人不存在，请刷新后重试"
	}
	return ""
}

func RSAEncode(str string) ([]byte, error) {
	origData := []byte(str)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RSADecode(ciphertext []byte) (string, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQC96+Z+ice9f/lBy+k4NBO5dJxrCk5R21YdGQLdz2Af9HJGYf75
97tQb04QyH0EQUah/oF6GM6iZxlOIaSspowK2StT44UtKao5SsQozJ8QFD6KXVwj
fgwHaOu3IKuAgBrclyQzRf/pOI57Jc8Y9Rv6WSaTReV/tCQoSADJ633l/wIDAQAB
AoGAGbaboVw0H9L4w1DBRau/U+eW2eMuUWTZ1tyxB6jxAcKNyjuwUtWYlb5MGnea
fX38+yfDDe3X5CMDSRHDAuEVqpKT91mdOBWX1qVicslHRwECnhLMVAMmpwCwTYCg
GXazjrwAl1zWylYGCY/V2X8ZhHLQLq5qhJWUJ6k9aVbps1kCQQDl90pazKVX+Ajh
IMtHXGXdb93JUSiBGUnfBpn+67D8rALH3DXOXDxbIE2C8wKBcgHCJ/SqapBIMsnz
h8jRogwLAkEA02wRrqxOYY1eG2s3GR91NJi/hZOgfufFuQAzLS+bnXxBFQjwWzdV
IfiP90vgwrAhJUW5c22xPmTL9F2xqcxSXQJBAOAKE75yMYOKedwafvB+7B7XpVNE
Zhmf8X/+hnj8VelUC0F7IFBzO7nrtpgk+AP0dhIZqxt7xiUQlf9UAil5nhECQQCC
k1AobUrLfSAOFx2kaoVcwqomuZJ6TnMTW0hANBMMJN2dPDQWYgo2POnNdhOOqnEO
MA3leG3rdx1wAx3jHMoRAkEA3pBDh2ojloZDhr5bhktZEDL/itS8X6jX8Vs4g965
bsl4uaE4DEIRQ5nZ534MOB9W+b49dhGhLP50Ti41nHifxA==
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`)
