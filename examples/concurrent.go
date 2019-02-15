package examples

import (
	"fmt"
	"time"
)

func goRoutineExample() {
	fmt.Println("<go routine>")
	counter := func(label string) {
		for i := 0; i < 3; i++ {
			fmt.Println(label, ":", i)
		}
	}

	// Call synchronously
	counter("direct call")

	// Execute concurrently
	go counter("go routine")

	// Anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("go routine + anonymous function")

	fmt.Println("<enter> to continue")
	fmt.Scanln()
}

func multipleExample() {
	// Starts several concurrent routines
	fmt.Println("<multiple>")
	letters := "ABCDEFGHIJ"
	for i := 0; i < 10; i++ {
		go func(i int) {
			for _, c := range letters {
				go func(c rune) {
					fmt.Printf("%c%d ", c, i)
				}(c)
			}
		}(i)
	}
	fmt.Println("<enter> to continue")
	fmt.Scanln()
}

func channelExample() {
	fmt.Println("<channel>")
	messages := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		messages <- "done"
	}()

	// Will block until message is received
	msg := <-messages
	fmt.Println(msg)
}

func bufferedExample() {
	messages := make(chan string, 2)

	// Place messages on buffered channel
	messages <- "buffered"
	messages <- "channel"

	// Consumes messages on buffer
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func synchronizationExample() {

	worker := func(done chan bool) {
		fmt.Println("working...")
		time.Sleep(time.Second)
		fmt.Println("done")
		// Tell channel work is done
		done <- true
	}

	done := make(chan bool, 1)

	// Call function asynchronously
	go worker(done)

	// Wait for signal on channel
	<-done
}

func channelDirectionsExample() {

	// May only send messages into channel pings
	ping := func(pings chan<- string, msg string) {
		pings <- msg
	}

	// May only receive messages from channel pings and send into channel pongs
	pong := func(pings <-chan string, pongs chan<- string) {
		msg := <-pings
		pongs <- msg
	}

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message to ping")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func deadlockExample() {
	dead := make(chan bool)
	// since channel is not being fed anything
	// this will produce a deadlock error
	<-dead
}

func selectExample() {

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "two"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "three"
	}()

	for i := 3; i > 0; i-- {
		fmt.Println("Waiting on", i)
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case msg3 := <-c3:
			fmt.Println("received", msg3)
		}
	}

}

// ConcurrentExamples contains examples of using go routines
func ConcurrentExamples() {
	fmt.Println("\nConcurrency in go")
	fmt.Println("=================")

	/*
		goRoutineExample()
		multipleExample()
		channelExample()
		bufferedExample()
		synchronizationExample()
		channelDirectionsExample()
	*/
	selectExample()
}
