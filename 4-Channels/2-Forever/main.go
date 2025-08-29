package main

// Thread 1
func main() {
	fovever := make(chan bool)

	go func() { // Thread 2
		for i := 0; i < 10; i++ {
			println(i)
		}
		fovever <- true
	}()

	<-fovever
}
