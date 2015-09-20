package main

import (
	"testing"
)

func TestEncode(t *testing.T) {
	result := Encode(int8(0x10), "test*", []byte{0x1d, 0x62, 0x40, 0xc5, 0xbd, 0x46, 0xb2, 0xab, 0x1f, 0x9e, 0x49, 0x92, 0x8c, 0x6a, 0x48, 0x81})
	if result != "f6a754391044d163a1021956ffadd15c" {
		t.Errorf("fail, expect %s, get %s", "f6a754391044d163a1021956ffadd15c", result)
	}
}

func TestGetSalt(t *testing.T) {
    _, _, err := get_salt("chenguan14")
    if err != nil {
        t.Error(err)
    }
}
