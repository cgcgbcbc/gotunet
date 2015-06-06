package main

import (
	"testing"
)

func TestIsTUNetIPAddress(t *testing.T) {
	if isTUNetIPAddress("59.66.250.114") && isTUNetIPAddress("166.111.81.219") && isTUNetIPAddress("101.5.213.43") {
		return
	}
	t.Error("fail")
}
