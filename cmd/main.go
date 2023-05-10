package main

import (
	"context"
	"fmt"
	"kafka_template/internal/kafkaRW"
)

const (
	TopicName     = "test_topic"
	BrokerAddress = "localhost:9092"
)

func main() {
	brokers := []string{BrokerAddress}

	ctx, cancel := context.WithCancel(context.Background())
	msgCh := make(chan kafkaRW.KafkaMessage, 1)
	go func() {
		if err := kafkaRW.ConsumeMessage(ctx, brokers, TopicName, msgCh); err != nil {
			cancel()
		}
	}()
	go func() {
		for {
			select {
			case msg := <-msgCh:
				// send to outer API using goroutines
				fmt.Println(msg)
			default:
			}
		}
	}()
}
