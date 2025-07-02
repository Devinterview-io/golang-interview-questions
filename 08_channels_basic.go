// 08_channels_basic.go - Basic channels demonstration
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== CHANNELS BASIC DEMONSTRATION ===")

	// Unbuffered channels
	fmt.Println("--- UNBUFFERED CHANNELS ---")
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Hello from goroutine!"
	}()

	fmt.Println("Waiting for message...")
	message := <-ch
	fmt.Printf("Received: %s\n", message)

	// Buffered channels
	fmt.Println("\n--- BUFFERED CHANNELS ---")
	bufferedCh := make(chan int, 3)

	// Send values without blocking
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3

	fmt.Printf("Channel length: %d, capacity: %d\n", len(bufferedCh), cap(bufferedCh))

	// Receive values
	for i := 0; i < 3; i++ {
		value := <-bufferedCh
		fmt.Printf("Received: %d\n", value)
	}

	// Channel direction (send-only, receive-only)
	fmt.Println("\n--- CHANNEL DIRECTIONS ---")
	dataCh := make(chan string, 2)

	go sender(dataCh)
	go receiver(dataCh)

	time.Sleep(2 * time.Second)

	// Closing channels
	fmt.Println("\n--- CLOSING CHANNELS ---")
	numberCh := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numberCh <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(numberCh) // Close the channel when done
	}()

	// Receive until channel is closed
	for number := range numberCh {
		fmt.Printf("Received number: %d\n", number)
	}

	fmt.Println("Channel closed, range loop ended")

	// Check if channel is closed
	testCh := make(chan int, 1)
	testCh <- 42
	close(testCh)

	value, ok := <-testCh
	fmt.Printf("Value: %d, Channel open: %t\n", value, ok)

	value, ok = <-testCh
	fmt.Printf("Value: %d, Channel open: %t\n", value, ok)
}

// Function that only sends to channel
func sender(ch chan<- string) {
	messages := []string{"First", "Second", "Third"}
	for _, msg := range messages {
		ch <- msg
		time.Sleep(300 * time.Millisecond)
	}
	close(ch)
}

// Function that only receives from channel
func receiver(ch <-chan string) {
	for msg := range ch {
		fmt.Printf("Receiver got: %s\n", msg)
	}
	fmt.Println("Receiver finished")
}
