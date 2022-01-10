package main

import "fmt"

func main() {
	//step1: we will create a channel.
	channel1 := make(chan int)

	go func() {

		for i := 0; i < 6; i++ {
			// send iterator over channel
			channel1 <- i
		}
		close(channel1)
		//if we'll not close the we'll get "fatal error: all goroutines are asleep - deadlock!"
		//
	}()

	//  range over channel to recv values
	for v := range channel1 {
		fmt.Println(v)
	}

}
