package kafka

import (
	"context"
	"errors"
	"publisher/models"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/mock"
)

// ============================================================
// External Repository Mock

type ProductRepoMock struct {
	mock.Mock
}

func (p *ProductRepoMock) Store(ctx context.Context, product models.Product) error {
	args := p.Called(product)

	return args.Error(0)
}

func (p *ProductRepoMock) Update(ctx context.Context, product models.Product) error {
	args := p.Called(product)

	return args.Error(0)
}

func (p *ProductRepoMock) Delete(ctx context.Context, id int64) error {
	args := p.Called(id)

	return args.Error(0)
}

// ============================================================
// Internal Repository Mock

type mockKafkaConnection struct {
	mockWriteMessages MockWriteMessages
}

type MockWriteMessages func(ctx context.Context, msgs ...kafka.Message) error

func (mock mockKafkaConnection) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	if len(msgs) < 1 {
		return errors.New("Should contain message")
	}

	return mock.mockWriteMessages(ctx, msgs...)
}
