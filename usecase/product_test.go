package usecase

import (
	"context"
	"publisher/models"
	kafkaRepoMock "publisher/repository/kafka"
	"publisher/repository/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Parallel()

	pgRepo := new(postgres.ProductRepoMock)
	kafkaRepo := new(kafkaRepoMock.ProductRepoMock)
	uc := NewProductUsecase(pgRepo, kafkaRepo)

	pgRepo.On("Get", int64(1)).Return(models.Product{}, nil)

	_, err := uc.Get(context.Background(), 1)
	assert.NoError(t, err)

}

func TestStore(t *testing.T) {
	t.Parallel()

	pgRepo := new(postgres.ProductRepoMock)
	kafkaRepo := new(kafkaRepoMock.ProductRepoMock)
	uc := NewProductUsecase(pgRepo, kafkaRepo)

	var (
		product = models.Product{
			ID:       1,
			Name:     "baju",
			Price:    1000,
			Rating:   1,
			ImageURL: "http://google.com",
		}
	)

	pgRepo.On("Store", product).Return(1, nil)
	kafkaRepo.On("Store", product).Return(nil)

	err := uc.Store(context.Background(), product)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	pgRepo := new(postgres.ProductRepoMock)
	kafkaRepo := new(kafkaRepoMock.ProductRepoMock)
	uc := NewProductUsecase(pgRepo, kafkaRepo)

	var (
		product = models.Product{
			ID:       1,
			Name:     "baju",
			Price:    1000,
			Rating:   1,
			ImageURL: "http://google.com",
		}
	)

	pgRepo.On("Get", int64(product.ID)).Return(product, nil)
	pgRepo.On("Update", product).Return(nil)
	kafkaRepo.On("Update", product).Return(nil)

	err := uc.Update(context.Background(), product)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	t.Parallel()

	pgRepo := new(postgres.ProductRepoMock)
	kafkaRepo := new(kafkaRepoMock.ProductRepoMock)
	uc := NewProductUsecase(pgRepo, kafkaRepo)

	pgRepo.On("Delete", int64(1)).Return(nil)
	kafkaRepo.On("Delete", int64(1)).Return(nil)

	err := uc.Delete(context.Background(), 1)
	assert.NoError(t, err)

}
