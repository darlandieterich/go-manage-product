package repository

import (
	"context"
	"product_manager/domain/model"

	"github.com/google/uuid"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) (uuid.UUID, error)
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, uuid uuid.UUID) error
	FindById(ctx context.Context, uuid uuid.UUID) (*model.Product, error)
	ListAll(ctx context.Context) ([]*model.Product, error)
}
