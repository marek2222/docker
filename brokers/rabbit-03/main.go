package main

import (
	"fmt"
	"time"

	"github.com/marek2222/docker/brokers/rabbit-03/rabbit"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := rabbit.GetConn("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			time.Sleep(time.Second * 2)
			conn.Publish("test-key", []byte(`{"message":"test"}`))
		}
	}()

	err = conn.StartConsumer("test-queue", "test-key", handler, 2)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	<-forever
}

func handler(d amqp.Delivery) bool {
	if d.Body == nil {
		fmt.Println("Error, no message body!")
		return false
	}
	fmt.Println(string(d.Body))
	return true
}
