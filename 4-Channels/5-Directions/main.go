package main

import "fmt"

func recebe(nome string, hello chan<- string) {
	hello <- nome
}

func ler(data <-chan string) {
	fmt.Println(<-data)
}

// Thread 1
func main() {
	hello := make(chan string)
	go recebe("John", hello)
	ler(hello)
}
