package kafkaRW

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaMessage struct {
	Author      string `json:"author"`
	Commit      string `json:"commit"`
	Description string `json:"description"`
	ProjectKey  string `json:"projectKey"`
}

func ConsumeMessage(ctx context.Context, brokers []string, topic string, msgCh chan KafkaMessage) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
	})

	for {
		select {
		case <-ctx.Done():
			err := reader.Close()
			if err != nil {
				return err
			}
			return nil
		default:
			msg, err := reader.FetchMessage(ctx)
			if err != nil {
				return err
			}
			kafkaMsg, err := parseKafkaMsg(&msg)
			if err != nil {
				return err
			}
			err = reader.CommitMessages(ctx, msg)
			if err != nil {
				return err
			}
			msgCh <- kafkaMsg
			return nil
		}
	}
}

func parseKafkaMsg(msg *kafka.Message) (KafkaMessage, error) {
	var kafkaMsg KafkaMessage
	err := json.Unmarshal(msg.Value, &kafkaMsg)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
	} else {
		log.Printf("Received message: %+v", kafkaMsg)
	}
	return kafkaMsg, err
}
