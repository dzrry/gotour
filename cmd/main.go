package main

import (
	"fmt"
	"github.com/dzrry/gotour/task2"
	"github.com/dzrry/gotour/task3"
	"github.com/dzrry/gotour/task5"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func main() {
	// task 2 slices
	pic.Show(task2.Pic)
	// task 3 maps
	wc.Test(task3.WordCounts)
	//task 5 stringers
	hosts := map[string]task5.IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
