package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	log "github.com/sirupsen/logrus"
	"os"
)

func publish(data interface{}, topic string) {
	mechanism, _ := scram.Mechanism(scram.SHA256, os.Getenv("KAFKA_USERNAME"), os.Getenv("KAFKA_PASSWORD"))
	w := kafka.Writer{
		Addr:  kafka.TCP(os.Getenv("KAFKA_HOST")),
		Topic: topic,
		Transport: &kafka.Transport{
			SASL: mechanism,
			TLS:  &tls.Config{},
		},
	}

	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling games to JSON: %v\n", err)
	}

	err = w.WriteMessages(context.Background(), kafka.Message{Value: payload})
	if err != nil {
		fmt.Println(err)
	}

	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
}
