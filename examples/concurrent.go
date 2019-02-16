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

func timeoutExample() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}

func nonBlockingExample() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func botChatExample() {

	talk := make(chan string, 1)
	done := make(chan bool, 1)
	end := make(chan bool, 1)

	go func() {
		for i := 0; i < 5; i++ {
			talk <- fmt.Sprint("hey", i)
			time.Sleep(time.Second)
		}
		done <- true
	}()

	go func() {
		finish := false
		for !finish {
			select {
			case msg := <-talk:
				fmt.Println("listened: ", msg)
			case <-done:
				finish = true
			default:
				fmt.Println("waiting...")
				time.Sleep(333 * time.Millisecond)
			}
		}
		fmt.Println("bye")
		end <- true
	}()
	<-end
}

func closingExample() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

func rangeChannelExample() {
	queue := make(chan string, 3)

	for i := 0; i < 3; i++ {
		queue <- fmt.Sprint("item", i)
	}
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

// ConcurrentExamples contains examples of using go routines
func ConcurrentExamples() {
	fmt.Println("\nConcurrency in go")
	fmt.Println("=================")

	goRoutineExample()
	multipleExample()
	channelExample()
	bufferedExample()
	synchronizationExample()
	channelDirectionsExample()
	selectExample()
	// deadlockExample()
	timeoutExample()
	nonBlockingExample()
	botChatExample()
	closingExample()
	rangeChannelExample()
}
