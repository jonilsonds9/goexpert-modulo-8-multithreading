package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string)

	go func() { // Thread 2
		canal <- "Olá mundo!" // Esta cheio
	}()

	// Thread 1
	msg := <-canal
	fmt.Println(msg)
}
