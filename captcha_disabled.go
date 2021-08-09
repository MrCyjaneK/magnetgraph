// +build !captcha
package main

var capt = false

func checkCaptcha(id, input string) bool {
	return true
}
