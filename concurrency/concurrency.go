package concurrency

import (
	"fmt"
	"time"
)

func Test() {
	go greet("Nice to meet you!")
	go greet("How are you?")
	done := make(chan bool)
	go longTime("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!")
	<-done
}

func greet(input string) {
	fmt.Println(input)
}

func longTime(phrase string, done chan bool) {
	time.Sleep(3 * time.Second) //simulate a long time
	fmt.Println("Hello!", phrase)
	done <- true
}
