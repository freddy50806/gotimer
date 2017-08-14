package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_time(t *testing.T) {
	p := fmt.Println
	f := func() {
		printls("SS")
	}
	timer1 := time.AfterFunc(3*time.Second, f)
	// We'll start by getting the current time.
	<-timer1.C
	now := time.Now()
	p(now)
}
