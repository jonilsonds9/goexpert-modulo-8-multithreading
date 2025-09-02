package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id      int
	content string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{int(i), "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{int(i), "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg1 := <-c1: // rabbitmq
			fmt.Printf("Received from RabbitMQ: ID: %d - %s\n", msg1.id, msg1.content)

		case msg2 := <-c2: //Kafka
			fmt.Printf("Received from Kafka: ID: %d - %s\n", msg2.id, msg2.content)

		case <-time.After(time.Second * 3):
			println("Timeout")
			//default:
			//	println("No communication")
		}
	}
}
