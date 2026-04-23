package concurrency

import (
	"fmt"
	"time"
)

func Start() {
	done := make(chan bool) // channel
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)
	// fmt.Println(<-done)//waiting for result to come out of channel
	// for doneChan := range done {
	// 	fmt.Println(doneChan)
	// }
	for range done {
	}
}

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	done <- true
	close(done)
}
