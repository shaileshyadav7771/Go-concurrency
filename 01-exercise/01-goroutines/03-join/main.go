package main

import (
	"fmt"
	"sync"
)

func main() {
	//TODO: modify the program
	// to print the value as 1
	// deterministically.

	var data int

	var relax sync.WaitGroup

	relax.Add(1) //It means 1 wait It'll wait till this wait=0(decremented to 0)

	go func() {
		relax.Done() // Decrementing the value -1 so wait(Variable)=0
		data++
	}()

	relax.Wait() // here t will excecute when wait == 0

	fmt.Printf("the value of data is %v\n", data)
	fmt.Println("Done..")

}
