package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(ch1 chan<- string) {
	// send message on ch1
	//ch1 chan<- string we have defined send only channel
	//Let's sent a message on ch1
	ch1 <- "Shailesh-9920886044"
	// here we can't receive data becz It's send only fn i.e <- ch1
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1  //ch1 <-chan string synatx for only recieve
	message_on_ch1 := <-ch1
	// send it on ch2 // ch2 chan<- string
	ch2 <- message_on_ch1

}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine genMsg and relayMsg

	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	// recv message on ch2
	v := <-ch2
	fmt.Println("The received message on Ch2 is : ", v)
}
