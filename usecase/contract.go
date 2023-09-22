package usecase

import (
	"context"
	"publisher/models"
)

type ProductUsecase interface {
	Get(ctx context.Context, id int64) (models.Product, error)
	Update(ctx context.Context, product models.Product) error
	Delete(ctx context.Context, id int64) error
	Store(ctx context.Context, product models.Product) error
}
