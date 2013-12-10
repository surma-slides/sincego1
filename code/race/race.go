package main

import (
	"fmt"
	"time"
)

var sharedVariable = 1

func worker() {
	for {
		sharedVariable += 1
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go worker()
	for {
		fmt.Printf("> %d\n", sharedVariable)
		time.Sleep(1 * time.Second)
	}
}
