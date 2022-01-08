package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call

	go fun("goroutine-10")

	// goroutine with anonymous function

	go func() {
		fun("goroutine-20")

	}()

	// goroutine with function value call

	fv := fun
	go fv("goroutine-3")

	// wait for goroutines to end
	fmt.Println("Waiting for our go routine to excecute.")
	time.Sleep(100 * time.Microsecond)

	fmt.Println("done..")
}
