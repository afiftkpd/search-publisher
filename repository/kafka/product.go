package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"publisher/models"

	"github.com/segmentio/kafka-go"
)

type writeMessages interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
}

type productRepository struct {
	Conn writeMessages
}

func NewProductRepository(conn writeMessages) ProductRepository {
	return &productRepository{conn}
}

func (p *productRepository) Store(ctx context.Context, product models.Product) error {
	convertedProduct, err := json.Marshal(product)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte("store"),
		Value: convertedProduct,
	}

	err = p.Conn.WriteMessages(context.Background(), msg)
	fmt.Println(err)
	return err
}

func (p *productRepository) Update(ctx context.Context, product models.Product) error {
	convertedProduct, err := json.Marshal(product)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte("update"),
		Value: convertedProduct,
	}

	err = p.Conn.WriteMessages(context.Background(), msg)
	return err
}

func (p *productRepository) Delete(ctx context.Context, id int64) error {
	msg := kafka.Message{
		Key:   []byte("delete"),
		Value: []byte(fmt.Sprintf("%d", id)),
	}

	err := p.Conn.WriteMessages(context.Background(), msg)
	return err
}
