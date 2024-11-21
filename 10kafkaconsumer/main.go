package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func topicSubscribtion(topic string, subChan chan string) {

}

func main() {

	topic := "test"

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	defer consumer.Close()

	err = consumer.Subscribe(topic, nil)

	if err != nil {
		panic(err)
	}

	println(
		"Subscribed",
	)

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			println("Message on", msg.TopicPartition.Partition, ":", string(msg.String()), string(msg.Value))
		} else {
			println("Consumer error:", err.Error())
		}
	}
}
