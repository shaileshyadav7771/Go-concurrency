package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	// TODO: implement timeout for recv on channel ch

	select {
	case m := <-ch:
		fmt.Println("message", m)

	case <-time.After(5 * time.Second):
		fmt.Println("Timeout...")
	}

}
