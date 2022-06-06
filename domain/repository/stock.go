package repository

import (
	"context"
	"product_manager/domain/model"

	"github.com/google/uuid"
)

type StockRepository interface {
	Create(ctx context.Context, stock *model.Stock) (uuid.UUID, error)
	Update(ctx context.Context, stock *model.Stock) error
	Delete(ctx context.Context, uuid uuid.UUID) error
	FindByProductID(ctx context.Context, productID uuid.UUID) (*model.Stock, error)
	FindAll(ctx context.Context, productID uuid.UUID) ([]*model.Stock, error)
}
