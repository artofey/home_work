package main

import (
	"log"
	"time"

	"github.com/beevik/ntp"
)

const (
	ntpHost = "0.beevik-ntp.pool.ntp.org"
)

func main() {
	beevikTime, err := ntp.Time(ntpHost)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("current time: %v\n", time.Now().Round(0))
	log.Printf("exact time: %v\n", beevikTime.Round(0))
}
