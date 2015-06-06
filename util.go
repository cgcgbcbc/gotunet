package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func read_response_body(resp *http.Response) (result string, err error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func getMacAddress() (mac string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	index := -1
	for i, addr := range addrs {
		if isTUNetIPAddress(addr.String()) {
			index = i
		}
	}
	if index != -1 {
		err = fmt.Errorf("no tsinghua network")
	}

	inter, err := net.InterfaceByIndex(index)
	if err != nil {
		return
	}
	mac = inter.HardwareAddr.String()
	return
}

func isTUNetIPAddress(ip string) bool {
	if strings.HasPrefix(ip, "59.66.") || strings.HasPrefix(ip, "166.111.") ||
		strings.HasPrefix(ip, "101.5.") || strings.HasPrefix(ip, "101.6.") ||
		strings.HasPrefix(ip, "183.173.") {
		return true
	}
	return false
}
