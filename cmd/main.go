package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

const (
	TopicName     = "test-topic"
	BrokerAddress = "localhost:9092"
	NumPartitions = 0
)

var ExampleProductions = [3]string{
	"A spectre is haunting Europe — the spectre of communism.",
	"All the powers of old Europe have entered into a holy alliance to exorcise this spectre:",
	"Pope and Tsar, Metternich and Guizot, French Radicals and German police-spies.",
}

func main() {
	if err := run(); err != nil && !errors.Is(err, context.Canceled) {
		log.Println(err)
		log.Fatal("stopping")
	}
}

func run() error {
	conn, err := kafka.DialLeader(context.Background(), "tcp", BrokerAddress, TopicName, NumPartitions)
	if err != nil {
		return err
	}
	defer func(c *kafka.Conn) {
		err := c.Close()
		if err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}(conn)
	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return err
	}
	for _, str := range ExampleProductions {
		_, err := conn.WriteMessages(
			kafka.Message{Value: []byte(str)})
		if err != nil {
			return fmt.Errorf("failed to write a message: %v", err)
		}
	}
	return nil
}
