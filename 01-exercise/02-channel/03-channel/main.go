package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 6)
	//here we have defined buffered value :) that is 6.

	go func() {

		//defer will excecute at last so we defining it first only in case if we forgot to define it.
		defer close(ch)

		// TODO: send all iterator values on channel without blocking
		for i := 0; i < 60000; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
