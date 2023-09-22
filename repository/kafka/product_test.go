package kafka

import (
	"context"
	"publisher/models"
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

func Test_Store(t *testing.T) {
	t.Parallel()

	var data = models.Product{
		ID:       1,
		Name:     "baju",
		Price:    100,
		Rating:   1,
		ImageURL: "google.com",
	}

	mockWriteMesasges := func(ctx context.Context, msgs ...kafka.Message) error {
		return nil
	}

	mock := mockKafkaConnection{mockWriteMessages: mockWriteMesasges}
	repo := NewProductRepository(mock)

	err := repo.Store(context.Background(), data)
	assert.NoError(t, err)
}

func Test_Update(t *testing.T) {
	t.Parallel()

	var data = models.Product{
		ID:       1,
		Name:     "baju",
		Price:    100,
		Rating:   1,
		ImageURL: "google.com",
	}

	mockWriteMesasges := func(ctx context.Context, msgs ...kafka.Message) error {
		return nil
	}

	mock := mockKafkaConnection{mockWriteMessages: mockWriteMesasges}
	repo := NewProductRepository(mock)

	err := repo.Update(context.Background(), data)
	assert.NoError(t, err)
}

func Test_Delete(t *testing.T) {
	t.Parallel()

	var data = models.Product{
		ID:       1,
		Name:     "baju",
		Price:    100,
		Rating:   1,
		ImageURL: "google.com",
	}

	mockWriteMesasges := func(ctx context.Context, msgs ...kafka.Message) error {
		return nil
	}

	mock := mockKafkaConnection{mockWriteMessages: mockWriteMesasges}
	repo := NewProductRepository(mock)

	err := repo.Delete(context.Background(), data.ID)
	assert.NoError(t, err)
}
