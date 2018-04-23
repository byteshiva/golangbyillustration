package main

import (
	"fmt"
	"time"
)

func f(a string, b string) {
	fmt.Printf("%s %s", a, b)
}
func makeMessage(s string, channel chan string) {
	for i := 0; i < 10; i++ {
		channel <- fmt.Sprintf("I wanted to tell you '%s' for the %dth time", s, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// go f("foo", "bar")
	// f("foo2", "bar2")
	// makeMessage("Hello world")
	channel := make(chan string) // create channel

	go makeMessage("Hello!", channel) // run this in a separate goroutine

	for i := 0; i < 10; i++ {
		fmt.Printf("Hey there, %s\n", <-channel) // read from channel and print out message
	}

	fmt.Println("Cool, that's all I wanted to say")
}
