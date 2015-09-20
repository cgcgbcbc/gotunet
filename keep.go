package main

import (
    "time"
    "log"
)

var (
    round = 5 * time.Minute
    not_online = "not_online"
)

func Keep(username string, password string) (err error) {
    for {
        status, err := GetStatus()
        if err != nil {
            log.Println(err)
        }
        log.Printf("status: %s", status)
        if err != nil || status == not_online {
            Login(username, password)
        }
        time.Sleep(round)
    }
    return
}
