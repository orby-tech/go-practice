package main

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func createTopic(topic string) {

	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
	})

	if err != nil {
		println("Failed to create topic:", err)

		panic(err)
	}

	defer admin.Close()

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	results, err := admin.CreateTopics(
		ctx,

		[]kafka.TopicSpecification{{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1}})

	if err != nil {
		println("Failed to create topic:", err)
		panic(err)
	}

	for _, result := range results {
		if result.Error.Code() != kafka.ErrNoError &&
			result.Error.Code() != kafka.ErrTopicAlreadyExists {
			println("Error creating topic:", result.Error.Error())
			panic(result.Error)

		} else {
			println("Topic", result.Topic, "created")
		}
	}

}

func topicSubscribtion(topic string, subChan chan string) {

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

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			println("Message on", msg.TopicPartition.Partition, ":", string(msg.String()))

			subChan <- string(msg.String())
		} else {
			println("Consumer error:", err)
		}
	}
}

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
	})

	if err != nil {
		panic(err)
	}

	defer producer.Close()

	deliveryChan := make(chan kafka.Event)

	topic := "test"

	for _, word := range []string{"Welcome", "to", "Kafka", "Golang"} {
		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}

		producer.Produce(message, deliveryChan)

		e := <-deliveryChan

		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			println("Delivery failed:", m.TopicPartition.Error)
			println("Failed message:", string(m.Value))
			println(e)
		} else {
			println("Delivered message to topic", *m.TopicPartition.Topic, "partition", m.TopicPartition.Partition, string(m.Value))
		}
	}

	close(deliveryChan)
}
