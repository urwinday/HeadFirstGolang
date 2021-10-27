package main

import (
	"fmt"
	"time"
)

func main1() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
}
