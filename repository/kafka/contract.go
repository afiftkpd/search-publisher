package kafka

import (
	"context"
	"publisher/models"
)

type ProductRepository interface {
	Store(ctx context.Context, product models.Product) error
	Update(ctx context.Context, product models.Product) error
	Delete(ctx context.Context, id int64) error
}
