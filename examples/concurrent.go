package examples

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
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

	// enqueue items
	for i := 0; i < 3; i++ {
		queue <- fmt.Sprint("item", i)
	}
	close(queue)

	// dequeue items
	for elem := range queue {
		fmt.Println(elem)
	}
}

func mySleep(d time.Duration) {
	// equivalent to time.Sleep(d)
	t := time.NewTimer(d)
	<-t.C
}

func timerExample() {
	timer1 := time.NewTimer(2 * time.Second)

	// blocks until 2s has passed
	<-timer1.C
	fmt.Println("Time 1 expired")

	mySleep(time.Millisecond * 500)

	// set a timer, but stop it before it's due
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

func tickerExample() {

	// send a new tick every 500ms in the channel
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		// handle ticks in separate thread
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// stop the ticker after 1600ms
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func workerPoolExample() {

	worker := func(id int, jobs <-chan int, results chan<- int) {
		for j := range jobs {
			fmt.Println("worker", id, "started job", j)
			time.Sleep(time.Second)
			fmt.Println("worker", id, "finished job", j)
			results <- j * 2
		}
	}

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	const numWorkers = 3
	const numJobs = 10

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Place jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// Collect results
	for r := 1; r <= numJobs; r++ {
		<-results
	}
}

func rateLimitExample() {

	// Enqueue 5 requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Limit handling a request every 400ms
	limiter := time.Tick(400 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Allow short bursts (3 max)
	burstyLimiter := make(chan time.Time, 3)

	// Leave 3 "ticks" in the channel
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Keep adding "ticks" to the channel (if possible, max=3 as above)
	go func() {
		for t := range time.Tick(400 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Enqueue requests
	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)

	// Handle requests limiting by the burstyLimiter channel allowin short bursts (3)
	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("burstyRequest", req, time.Now())
	}
}

func atomicExample() {

	var opsAtomic uint64
	var opsNonAtomic uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&opsAtomic, 1)
				opsNonAtomic++
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(5 * time.Second)

	opsAtomicFinal := atomic.LoadUint64(&opsAtomic)
	opsNonAtomicFinal := opsNonAtomic
	fmt.Println("opsAtomic (exact):", opsAtomicFinal)
	fmt.Println("opsNonAtomic (unsafe):", opsNonAtomicFinal)

}

func mutexExample() {

	// Mutually Exclusive (mutex) locking
	var mutex = &sync.Mutex{}

	// Application shared state
	var state = make(map[int]int)

	// Counters (updated atomically)
	var readOps, writeOps uint64

	// Start 100 readers and safely read from the state by locking a mutex
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)

				mutex.Lock()
				total += state[key]
				mutex.Unlock()

				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Start 10 writers and safely write to the state by locking using a mutex
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)

				mutex.Lock()
				state[key] = val
				mutex.Unlock()

				atomic.AddUint64(&writeOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let readers and writers work
	time.Sleep(time.Second)

	// Get updated counters
	fmt.Println("Reads: ", atomic.LoadUint64(&readOps))
	fmt.Println("Writes: ", atomic.LoadUint64(&writeOps))

	// Get current state (safelly through the lock!)
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

func statefulExample() {

	// Types used to pass the parameters to the stateful go routine

	type readOp struct {
		key  int
		resp chan int
	}

	type writeOp struct {
		key  int
		val  int
		resp chan bool
	}

	// Counters (incremented atomically)
	var readOps uint64
	var writeOps uint64

	// Queue for operations
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		// Private state for go routine
		var state = make(map[int]int)
		// Keep receiving messages in the read and write channels
		for {
			select {
			case read := <-reads:
				// Send response in the "response channel"
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Start 100 readers
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// Create message and enqueue it in "reads"
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int, 1)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				// Create message and enqueue it in "writes"
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool, 1)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Run for a whole second
	time.Sleep(time.Second)

	// Results
	fmt.Println("Reads:", atomic.LoadUint64(&readOps))
	fmt.Println("Writes:", atomic.LoadUint64(&writeOps))
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
		selectExample()
		// deadlockExample()
		timeoutExample()
		nonBlockingExample()
		botChatExample()
		closingExample()
		rangeChannelExample()
		timerExample()
		tickerExample()
		workerPoolExample()
		rateLimitExample()
		atomicExample()
		mutexExample()
	*/
	statefulExample()
}
