package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func Encode(v1 int8, password string, challenge []byte) (result string) {
	passwordMd5 := md5.Sum([]byte(password))
	temp := fmt.Sprintf("%s%s%s", string(v1), hex.EncodeToString(passwordMd5[:]), string(challenge))
	md5sum := md5.Sum([]byte(temp))
	result = hex.EncodeToString(md5sum[:])
	return
}

func Login(username string, password string) (result string, err error) {
	v1, challenge, err := get_salt(username)
	if err != nil {
		return
	}
	epwd := Encode(v1, password, challenge)
	return do_login(username, epwd)
}

func Logout() (result string, err error) {
	resp, err := http.PostForm("http://166.111.8.120:3333/cgi-bin/do_logout", nil)
	if err != nil {
		return
	}
	return read_response_body(resp)
}

func do_login(username string, epwd string) (result string, err error) {
	mac, err := getMacAddress()
	if err != nil {
		return
	}
    url := "http://166.111.204.120:69/cgi-bin/srun_portal"
	data := fmt.Sprintf("action=login&username=%s&drop=0&pop=0&type=2&n=117&mbytes=0&minutes=0&ac_id=1&password=%s&chap=1&mac=%s", username, epwd, mac)
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(data))
	if err != nil {
		return
	}
	return read_response_body(resp)
}

func build_salt_message(username string) (message []byte) {
	message = make([]byte, 56)
	copy(message[0:16], []byte{0x9c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	copy(message[16:], username)
	return
}

func get_salt(username string) (user_id int8, challenge []byte, err error) {
	message := build_salt_message(username)

	raddr, err := net.ResolveUDPAddr("udp", "166.111.204.120:3335")
	if err != nil {
		return
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	defer conn.Close()
	if err != nil {
		return
	}

	msg := make(chan string)
	e := make(chan error)

	go func() {
		var buf [64]byte
		for {
			_, err = conn.Read(buf[0:])
			msg <- string(buf[0:])
			e <- err
			break
		}
	}()

	_, err = conn.Write(message)
	if err != nil {
		return
	}
	response := []byte(<-msg)
	err = <-e
	if err != nil {
		return
	}
	user_id = int8(binary.BigEndian.Uint64(response[8:16]) >> 56)
	challenge = response[16:32]
	return
}
