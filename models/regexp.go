package models

import (
	"regexp"
	"strconv"
	"time"
)

const (
	pattern_string    = "^\\w{6,18}$"
	pattern_email     = "^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+.[A-Za-z]{2,4}$"
	pattern_birth     = "^19\\d{2}((0[1-9])|(1[0-2]))(0[1-9]|[12]\\d|3[01])$"
	pattern_odds      = "^\\d+(.\\d+)?$"
	pattern_score     = "^-?(\\d)+.5$"
	pattern_starttime = "^\\d{4}-\\d{2}-\\d{2} (\\d{2}:){2}\\d{2}$"
	pattern_betamount = "^\\d+(.\\d{1,8})?$"
	pattern_result    = "^\\d+:\\d+$"
	pattern_gid       = "^\\d+"
)

func ValidString(s string) bool {
	r := []rune(s)
	if !(len(r) <= 18 && len(r) >= 6 && len(r) == len(s)) {
		return false
	}
	if ok, _ := regexp.MatchString(pattern_string, s); !ok {
		return false
	}
	return true
}

func ValidEmail(s string) bool {
	if ok, _ := regexp.MatchString(pattern_email, s); !ok {
		return false
	}
	return true
}

func ValidBirth(s string) bool {
	if ok, _ := regexp.MatchString(pattern_birth, s); !ok {
		return false
	}
	return true
}

func ValidOdds(s string) bool {
	if ok, _ := regexp.MatchString(pattern_odds, s); !ok {
		return false
	}
	return true
}

func ValidScore(s string) bool {
	if ok, _ := regexp.MatchString(pattern_score, s); !ok {
		return false
	}
	return true
}

func ValidStarttime(s string) bool {
	if ok, _ := regexp.MatchString(pattern_starttime, s); !ok {
		return false
	}

	timeStart, _ := time.Parse(layout, s)
	timeNow, _ := time.Parse(layout, time.Now().String())
	if timeStart.Sub(timeNow).Hours() < 6 {
		return false
	}
	return true
}

func ValidBetamount(s string) bool {
	if ok, _ := regexp.MatchString(pattern_betamount, s); !ok {
		return false
	}

	amount, _ := strconv.ParseFloat(s, 64)
	if amount <= 0.0001 {
		return false
	}

	return true
}

func ValidResult(s string) bool {
	if ok, _ := regexp.MatchString(pattern_result, s); !ok {
		return false
	}

	return true
}

func ValidGid(s string) bool {
	if ok, _ := regexp.MatchString(pattern_gid, s); !ok {
		return false
	}

	return true
}
