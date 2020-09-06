package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func checkErr(err error, msg string ) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

func handleWork(messageBody []byte) error {
	// TODO : write worker logic here. this is just a sample
	
	var msgMap map[string]json.RawMessage
	err := json.Unmarshal(messageBody, &msgMap)
	checkErr(err, "error in json decoding msg")

	log.Println("sleeping for 2 sec ...")
	time.Sleep(2 * time.Second)

	return nil
}

func main() {


	AMQP_CONNECTION_URL := os.Getenv("AMQP_CONNECTION_URL")
	log.Println("starting worker, amqp connection url : ",AMQP_CONNECTION_URL) // todo : remove this

	conn, err := amqp.Dial(AMQP_CONNECTION_URL)
	checkErr(err, "Can't connect to broker")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	checkErr(err, "Can't create a amqpChannel")
	defer amqpChannel.Close()

	// https://godoc.org/github.com/streadway/amqp#Channel.QueueDeclare
	queue, err := amqpChannel.QueueDeclare("test_queue", true, false, false, false, nil)
	checkErr(err, "Could not declare `test_queue` queue")

	// https://godoc.org/github.com/streadway/amqp#Channel.Qos
	err = amqpChannel.Qos(1, 0, false)
	checkErr(err, "Could not configure QoS")

	// https://godoc.org/github.com/streadway/amqp#Channel.Consume
	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	checkErr(err, "Could not register consumer")

	stopChan := make(chan bool)

	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			if err := handleWork(d.Body) ; err != nil {
				// https://godoc.org/github.com/streadway/amqp#Delivery.Nack
				nackErr := d.Nack(false, false)
				checkErr(nackErr, "")
				log.Printf("handleWork returned error, Nack-ed message: %s", d.Body)
			} else {
				// https://godoc.org/github.com/streadway/amqp#Delivery.Ack
				ackErr := d.Ack(false)
				checkErr(ackErr, "")
				log.Printf("handleWork successful, Ack-ed message: %s", d.Body)
			}

		}
	}()

	<-stopChan
}

