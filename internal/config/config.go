package config

import (
	"github.com/caarlos0/env/v8"
	"time"
)

type Config struct {
	UserKey            string        `env:"USER_KEY"  envDefault:""`
	RequestURL         string        `env:"REQUEST_URL"  envDefault:"https://jira.sbercloud.dev/rest/api/2/issue"`
	ProjectKey         string        `env:"PROJECT_KEY"  envDefault:"INCIDENT"`
	IssueType          string        `env:"ISSUE_TYPE"  envDefault:"10002"` //10002
	RequestSendTimeout time.Duration `env:"REQUEST_SEND_TIMEOUT"  envDefault:"2s"`
	// Kafka содержит настройки для работы с Kafka.
	Kafka struct {
		// Brokers является готовой строкой подключения к брокерам
		// в формате, установленном выбранным брокером (kafkaRW).
		//
		// Предполагается строка следующего вида: "localhost:9092,localhost:9093".
		Brokers       []string `env:"KAFKA_BROKERS,notEmpty" envSeparator:","  envDefault:"localhost:9092"`
		User          string   `env:"KAFKA_USER,notEmpty" envDefault:"user"`
		Password      string   `env:"KAFKA_PASSWORD, notEmpty" envDefault:"bitnami"`
		FirewallTopic string   `env:"KAFKA_FIREWALL_TOPIC,notEmpty" envSeparator:","  envDefault:"firewall"`
	}
}

// Load возвращает настроенную конфигурацию сервиса.
func Load() (*Config, error) {
	c := Config{}
	if err := env.Parse(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
