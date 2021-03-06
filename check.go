package main

import (
	"net/http"
	"strings"
)

func CheckOnline() (result string, err error) {
	url := "http://166.111.8.120/cgi-bin/do_login"
	data := "action=check_online"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return
	}
	return read_response_body(resp)
}

func GetStatus() (result string, err error) {
	url := "http://166.111.204.120:69/cgi-bin/rad_user_info"
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	return read_response_body(resp)
}
