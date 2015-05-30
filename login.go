package main

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
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

func read_response_body(resp *http.Response) (result string, err error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func do_login(username string, epwd string) (result string, err error) {
	resp, err := http.PostForm("http://166.111.8.120:3333/cgi-bin/do_login",
		url.Values{"username": {username}, "type": {"2"}, "password": {epwd}, "chap": {"1"}})
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

	raddr, err := net.ResolveUDPAddr("udp", "166.111.8.120:3335")
	if err != nil {
		return
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return
	}

	defer conn.Close()

	msg := make(chan string)
	e := make(chan error)

	go func() {
		var buf [48]byte
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
