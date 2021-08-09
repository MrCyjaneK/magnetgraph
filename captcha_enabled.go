// +build captcha

package main

import "strings"

var capt = true

func checkCaptcha(id, input string) bool {
	if strings.EqualFold(captchas[id], input) && id != "" && captchas[id] != "" {
		delete(captchas, id)
		delete(timec, id)
		return true
	}
	return false
}
