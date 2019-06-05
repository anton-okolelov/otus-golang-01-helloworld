package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func main() {
	preciseTime, err := ntp.Time("time.google.com")
	if err == nil {
		fmt.Printf("Current Local time is   %s\n", time.Now())
		fmt.Printf("Current precise time is %s\n", preciseTime)
	} else {
		fmt.Printf("Error: %s", err);
	}
}
