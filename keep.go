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
        log.Printf("status: %s\n", status)
        if err != nil || status == not_online {
            log.Println("try login")
            result, err := Login(username, password)
            log.Printf("login result: %s\n", result)
            if err != nil {
                log.Println("login failed, retry after a while")
                log.Println(err)
            }
        }
        time.Sleep(round)
    }
    return
}
