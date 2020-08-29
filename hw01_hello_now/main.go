package main

import (
	"fmt"
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
	fmt.Printf("current time: %v\n", time.Now().Round(0))
	fmt.Printf("exact time: %v\n", beevikTime.Round(0))
}
