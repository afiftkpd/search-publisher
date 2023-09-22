package postgres

import (
	"context"
	"publisher/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	mock.Mock
}

func (p *ProductRepoMock) Store(ctx context.Context, product models.Product) (int, error) {
	args := p.Called(product)

	return args.Get(0).(int), args.Error(1)
}

func (p *ProductRepoMock) Delete(ctx context.Context, id int64) error {
	args := p.Called(id)

	return args.Error(0)
}

func (p *ProductRepoMock) Update(ctx context.Context, product models.Product) error {
	args := p.Called(product)

	return args.Error(0)
}

func (p *ProductRepoMock) Get(ctx context.Context, id int64) (models.Product, error) {
	args := p.Called(id)

	return args.Get(0).(models.Product), args.Error(1)
}

func (p *ProductRepoMock) Coba(ctx context.Context, product models.Product) (models.Product, error) {
	args := p.Called(product)

	return args.Get(0).(models.Product), args.Error(1)
}
