package main

func main() {
	ch := make(chan string, 2) // Buffered channel with capacity of 2
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
