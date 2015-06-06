package main

import (
	"net/http"
	"strings"
)

func CheckOnline() (result string, err error) {
	url := "http://net.tsinghua.edu.cn/cgi-bin/do_login"
	data := "action=check_online"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return
	}
	return read_response_body(resp)
}
