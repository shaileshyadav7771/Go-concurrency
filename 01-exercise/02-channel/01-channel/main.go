package main

import "fmt"

func main() {

	//step1: we will create a channel.
	ch := make(chan int)

	go func(a, b int) {
		c := a + b

		//step2:now let's send value of c variable from our go rountine to channel
		ch <- c

	}(1, 2)
	// TODO: get the value computed from goroutine

	//step3: We need to receive the value
	r := <-ch
	fmt.Printf("computed value %v\n", r)

	//Note: Here we are having our Go routine which is doing some computation and we want it's
	//result inside our main routine without of Sharing memory.
}
