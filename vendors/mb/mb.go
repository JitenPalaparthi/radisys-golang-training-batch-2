package mb

import (
	"context"
	"log"
	"vendors/models"

	"github.com/segmentio/kafka-go"
)

type DATA interface {
	models.Vendor
}

func Publish(conn []string, topic string, key []byte, value []byte) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: conn,
		Topic:   topic,
		// assign the logger to the writer
		//Logger: mb.Logger,
	})
	defer w.Close()
	return w.WriteMessages(context.TODO(), kafka.Message{Key: key, Value: value})
}

func Subscribe(conn []string, topic string) <-chan []byte {
	chanMsg := make(chan []byte)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: conn,
		Topic:   topic,
		// assign the logger to the writer
		//Logger: mb.Logger,
	})
	go func() {
		for {
			msg, err := r.ReadMessage(context.TODO())
			if err != nil {
				log.Println(err)
			} else {
				//fmt.Println(msg.Value)
				chanMsg <- msg.Value
			}
		}
	}()
	return chanMsg
}
