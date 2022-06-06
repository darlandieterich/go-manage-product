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
	FindByCode(ctx context.Context, code string) (*model.Product, error)
	FindAll(ctx context.Context) ([]*model.Product, error)
}
