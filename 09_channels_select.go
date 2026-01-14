// 09_channels_select.go - Select statement with channels
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== CHANNELS SELECT DEMONSTRATION ===")

	// Basic select statement
	fmt.Println("--- BASIC SELECT ---")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	// Select will choose the first channel that's ready
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("Received from ch2: %s\n", msg2)
		}
	}

	// Select with default case
	fmt.Println("\n--- SELECT WITH DEFAULT ---")
	ch3 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch3 <- "Delayed message"
	}()

	for i := 0; i < 5; i++ {
		select {
		case msg := <-ch3:
			fmt.Printf("Received: %s\n", msg)
		default:
			fmt.Printf("No message available (attempt %d)\n", i+1)
			time.Sleep(500 * time.Millisecond)
		}
	}

	// Select with timeout
	fmt.Println("\n--- SELECT WITH TIMEOUT ---")
	slowCh := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		slowCh <- "This message is too slow"
	}()

	select {
	case msg := <-slowCh:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! No message received within 1 second")
	}

	// Multiple channel communication
	fmt.Println("\n--- MULTIPLE CHANNEL COMMUNICATION ---")
	numbers := make(chan int)
	strings := make(chan string)
	done := make(chan bool)

	// Producer goroutines
	go func() {
		for i := 1; i <= 3; i++ {
			numbers <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(numbers)
	}()

	go func() {
		words := []string{"apple", "banana", "cherry"}
		for _, word := range words {
			strings <- word
			time.Sleep(400 * time.Millisecond)
		}
		close(strings)
	}()

	// Consumer
	go func() {
		numbersOpen := true
		stringsOpen := true

		for numbersOpen || stringsOpen {
			select {
			case num, ok := <-numbers:
				if ok {
					fmt.Printf("Received number: %d\n", num)
				} else {
					fmt.Println("Numbers channel closed")
					numbersOpen = false
				}
			case str, ok := <-strings:
				if ok {
					fmt.Printf("Received string: %s\n", str)
				} else {
					fmt.Println("Strings channel closed")
					stringsOpen = false
				}
			}
		}
		done <- true
	}()

	<-done
	fmt.Println("All channels processed")
}
