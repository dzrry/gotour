package main

import (
	"github.com/dzrry/gotour/task2"
	"github.com/dzrry/gotour/task3"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func main() {
	// task 2 slices
	pic.Show(task2.Pic)
	// task 3 maps
	wc.Test(task3.WordCounts)
}
