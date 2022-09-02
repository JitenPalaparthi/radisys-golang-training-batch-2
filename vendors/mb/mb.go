package mb

import (
	"context"
	"strconv"
	"vendors/models"

	"github.com/segmentio/kafka-go"
)

type DATA interface {
	models.Vendor
}

func Publish[T DATA](conn []string, topic string, data T) error {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: conn,
		Topic:   topic,
		// assign the logger to the writer
		//Logger: mb.Logger,
	})
	val, err := data.ToBytes()
	if err != nil {
		return err
	}
	return w.WriteMessages(context.TODO(), kafka.Message{Key: []byte(strconv.Itoa(data.ID)), Value: val})

}
