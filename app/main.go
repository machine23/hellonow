package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

const ntpServer = "pool.ntp.org"

func main() {
	curTime, err := ntp.Time(ntpServer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current time: %v\n", curTime)
}
