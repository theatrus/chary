package main

import (
	"fmt"
	"github.com/theatrus/chary/fitters"
	"time"
)

func main() {
	fit, _ := fitters.NewWindow(time.Duration(60)*time.Second, time.Duration(1)*time.Second)
	fmt.Println("hi")
	fmt.Println(fit.Buckets())
}
