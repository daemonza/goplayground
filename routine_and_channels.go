package main

import (
	"fmt"
)

func blocking(message string) {
	for a := 1; a <= 5; a++ {
		fmt.Println(message)
	}
}

func async(message string) {
	for a := 1; a <= 5; a++ {
		fmt.Println(a, ":", message)
	}
}

func asyncWorker(myChannel chan string) {
	for a := 1; a <= 5; a++ {
		fmt.Println("doing some work in a goroutine")
	}

	// Send message on myChannel that I am done.
	myChannel <- "done"

}

func asyncReceiver(myChannel chan string) {
	// message is whatever is in the myChannel
	message := <-myChannel
	fmt.Println("I got : ", message)

}

func main() {

	// Create a channel type.
	// Format is : variable name, the type which is a channel
	// and then string which is the type we pass over the channel
	var myChannel chan string = make(chan string)

	// Example of the blocking function blocking the
	// execution to our goroutine that execute async()
	blocking("I am a normal blocking function")
	go async("async1")

	// Example of the blocking function after the goroutine
	go async("async2")
	blocking("I got executed after async1 and async2")

	// Two goroutine functions showing how channels work.
	// asyncWorker does some work and when completed sends a
	// message on myChannel which asyncReceiver listen on and
	// then prints out.
	go asyncWorker(myChannel)
	go asyncReceiver(myChannel)

	var input string
	fmt.Scanln(&input)

	fmt.Println("done")

}
