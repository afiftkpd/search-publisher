package usecase

import (
	"context"
	"errors"
	"publisher/models"
	kafkaR "publisher/repository/kafka"
	"publisher/repository/postgres"
)

type productUsecase struct {
	PostgresRepo postgres.ProductRepository
	KafkaRepo    kafkaR.ProductRepository
}

func NewProductUsecase(postgresRepo postgres.ProductRepository, kafkaRepo kafkaR.ProductRepository) ProductUsecase {
	return &productUsecase{postgresRepo, kafkaRepo}
}

func (p *productUsecase) Get(ctx context.Context, id int64) (models.Product, error) {
	result, err := p.PostgresRepo.Get(ctx, id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (p *productUsecase) Update(ctx context.Context, product models.Product) error {
	result, err := p.Get(ctx, product.ID)
	if err != nil {
		return err
	}

	if result.ID == 0 {
		return errors.New("id not found")
	}

	return p.KafkaRepo.Update(ctx, product)
}

func (p *productUsecase) Delete(ctx context.Context, id int64) error {
	err := p.PostgresRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return p.KafkaRepo.Delete(ctx, id)
}

func (p *productUsecase) Store(ctx context.Context, product models.Product) error {
	id, err := p.PostgresRepo.Store(ctx, product)
	if err != nil {
		return err
	}

	product.ID = int64(id)
	return p.KafkaRepo.Store(ctx, product)
}
