package main

import (
	"fmt"
	"github.com/dzrry/gotour/task2"
	"github.com/dzrry/gotour/task3"
	"github.com/dzrry/gotour/task5"
	"github.com/dzrry/gotour/task7"
	"github.com/dzrry/gotour/task8"
	"github.com/dzrry/gotour/task9"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
	"golang.org/x/tour/wc"
	"io"
	"os"
	"strings"
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
	//task 7 reader
	reader.Validate(task7.MyReader{})
	//task 8 rot13
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := task8.Rot13Reader{s}
	io.Copy(os.Stdout, &r)
	//task 9 image
	i := &task9.Image{
		Width:  100,
		Height: 100,
		Clr:    100,
	}
	pic.ShowImage(i)
}
